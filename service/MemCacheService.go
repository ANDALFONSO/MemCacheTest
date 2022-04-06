package service

import (
	"encoding/json"
	"fmt"
	"memcache/entity"
	"os"
	"strconv"
	"strings"

	"github.com/mercadolibre/go-meli-toolkit/gomemcached"
)

const (
	MEM_CACHE = "shipping-wdm-api"
)

type IMemCacheService interface {
	CreateKey(key string, personalData *entity.PersonalData) string
	DeleteKey(key string) string
	GetKey(key string) *entity.PersonalData
}

type MemCacheService struct {
	client gomemcached.Client
}

func NewMemCache() IMemCacheService {
	endPoindNodes := os.Getenv("CACHE_CACHE_TEST_CONFIG_ENDPOINT")
	fmt.Println("endPoindNode:", endPoindNodes)
	nodes := strings.Split(os.Getenv("CACHE_CACHE_TEST_NODES_ENDPOINT"), ",")
	fmt.Println("Nodes:", nodes)

	gomemcached.RegisterCluster(MEM_CACHE)
	client, _ := gomemcached.NewClient(MEM_CACHE)

	return &MemCacheService{
		client: client,
	}
}

func (m *MemCacheService) CreateKey(key string, personalData *entity.PersonalData) string {
	data, _ := json.Marshal(personalData)
	expiration_item, _ := strconv.Atoi(os.Getenv("EXPIRATION_ITEM"))
	err := m.client.Set(key, data, int32(expiration_item))
	if err != nil {
		return fmt.Sprintf("error creating key %v", key)
	}
	return fmt.Sprintf("creation key %v ok", key)
}

func (m *MemCacheService) DeleteKey(key string) string {
	err := m.client.Delete(key)
	if err != nil {
		return fmt.Sprintf("Error deleting key %v", key)
	}
	return fmt.Sprintf("delete key %v ok", key)
}

func (m *MemCacheService) GetKey(key string) *entity.PersonalData {
	byte, err := m.client.Get(key)
	if err != nil {
		fmt.Printf("error key %v not found: %v", key, err)
		return &entity.PersonalData{}
	}

	var personalData entity.PersonalData
	err = json.Unmarshal(byte, &personalData)
	if err != nil {
		fmt.Println("Error casting personal data. ", err)
		return &entity.PersonalData{}
	}
	return &personalData
}
