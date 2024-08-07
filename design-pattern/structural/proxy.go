package structural

import "fmt"

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

type FeatureAccess interface {
	accessFeature(userId string) string
}

type FeatureAccessService struct{}

func NewFeatureAccessService() *FeatureAccessService {
	return &FeatureAccessService{}
}

func (f *FeatureAccessService) accessFeature(userId string) string {
	return fmt.Sprintf("User %s has accessed the premium feature", userId)
}

type SubscriptionProxy struct {
	featureAccessService *FeatureAccessService
	subscribedUsers      []string
}

func NewSubscriptionProxy(featureAccessService *FeatureAccessService) *SubscriptionProxy {
	return &SubscriptionProxy{
		featureAccessService: featureAccessService,
		subscribedUsers:      []string{},
	}
}

func (f *SubscriptionProxy) accessFeature(userId string) string {
	if !Contains(f.subscribedUsers, userId) {
		return fmt.Sprintf("User %s is not subscribed and cannot access premium feature", userId)
	}
	return f.featureAccessService.accessFeature(userId)
}

func (f *SubscriptionProxy) subscribeUser(userId string) {
	f.subscribedUsers = append(f.subscribedUsers, userId)
}

func clientCodeProxy(featureAccess FeatureAccess, userId string) {
	res := featureAccess.accessFeature(userId)
	fmt.Println(res)
}

func runProxy() {
	realService := NewFeatureAccessService()
	proxy := NewSubscriptionProxy(realService)

	proxy.subscribeUser("user123")

	clientCodeProxy(proxy, "user123")
	clientCodeProxy(proxy, "user124")

}
