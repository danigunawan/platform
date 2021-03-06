/*
 * Copyright (c) 2019. Pandranki Global Private Limited
 */

package models

func GetCoreServiceScopes() []string {
	serviceScopes := []string{
		"RentalPackage:Read",
		"RentalPackage:List",
		"Currency:Read",
		"Currency:List",
		"WebhookLog:Read",
		"WebhookLog:List",
		"StoreVehicleType:Read",
		"StoreVehicleType:List",
		"Webhook:Read",
		"Webhook:List",
		"OAuthApplication:Read",
		"OAuthApplication:List",
		"AdminDashboard:Read",
		"AdminDashboard:List",
		"SEOSetting:Read",
		"SEOSetting:List",
		"MarketSetting:Read",
		"MarketSetting:List",
		"JobTimeVariance:Read",
		"JobTimeVariance:List",
		"JobRequestAcceptanceReport:Read",
		"JobRequestAcceptanceReport:List",
		"Job:Read",
		"Job:List",
		"ProviderLogReport:Read",
		"ProviderLogReport:List",
		"UserWalletReport:Read",
		"UserWalletReport:List",
		"StoreReview:Read",
		"StoreReview:List",
		"CancelledReport:Read",
		"CancelledReport:List",
		"ProviderPaymentReport:Read",
		"ProviderPaymentReport:List",
		"StorePaymentReport:Read",
		"StorePaymentReport:List",
		"AdminReport:Read",
		"AdminReport:List",
		"WineDeliveryLabel:Read",
		"WineDeliveryLabel:List",
		"GroceryDeliveryLabel:Read",
		"GroceryDeliveryLabel:List",
		"FoodDeliveryLabel:Read",
		"FoodDeliveryLabel:List",
		"GeneralLabel:Read",
		"GeneralLabel:List",
		"AirportSurcharge:Read",
		"AirportSurcharge:List",
		"LocationWiseFare:Read",
		"LocationWiseFare:List",
		"DeliveryCharge:Read",
		"DeliveryCharge:List",
		"DeclineAlert:Read",
		"DeclineAlert:List",
		"HelpCategory:Read",
		"HelpCategory:List",
		"HelpDetail:Read",
		"HelpDetail:List",
		"FAQCategory:Read",
		"FAQCategory:List",
		"FAQ:Read",
		"FAQ:List",
		"NewsletterSubscriber:Read",
		"NewsletterSubscriber:List",
		"EnterpriseAccount:Read",
		"EnterpriseAccount:List",
		"BusinessTripReason:Read",
		"BusinessTripReason:List",
		"RideProfileType:Read",
		"RideProfileType:List",
		"VisitLocation:Read",
		"VisitLocation:List",
		"VehicleModel:Read",
		"VehicleModel:List",
		"VehicleMake:Read",
		"VehicleMake:List",
		"SMSTemplate:Read",
		"SMSTemplate:List",
		"EmailTemplate:Read",
		"EmailTemplate:List",
		"GeoFenceRestrictedArea:Read",
		"GeoFenceRestrictedArea:List",
		"GeoFenceLocation:Read",
		"GeoFenceLocation:List",
		"DeliveryChargesUtility:Read",
		"DeliveryChargesUtility:List",
		"OrderStatusUtility:Read",
		"OrderStatusUtility:List",
		"Order:Read",
		"Order:List",
		"StoreItemType:Read",
		"StoreItemType:List",
		"StoreItem:Read",
		"StoreItem:List",
		"StoreItemCategory:Read",
		"StoreItemCategory:List",
		"DeliveryVehicleType:Read",
		"DeliveryVehicleType:List",
		"Store:Read",
		"Store:List",
		"AdvertisementBanner:Read",
		"AdvertisementBanner:List",
		"View:Read",
		"View:List",
		"CancelReason:Read",
		"CancelReason:List",
		"PackageType:Read",
		"PackageType:List",
		"Page:Read",
		"Page:List",
		"User:Read",
		"User:List",
		"Review:Read",
		"Review:List",
		"Coupon:Read",
		"Coupon:List",
		"ServiceType:Read",
		"ServiceType:List",
		"ServiceSubCategory:Read",
		"ServiceSubCategory:List",
		"Service:Read",
		"Service:List",
		"RequiredDocument:Read",
		"RequiredDocument:List",
		"ServiceProvider:Read",
		"ServiceProvider:List",
		"ServiceCompany:Read",
		"ServiceCompany:List",
		"IAMGroup:Read",
		"IAMGroup:List",
		"MarketStatistics:Read",
		"MarketStatistics:List",
		"AppInstallation:Read",
		"AppInstallation:List",
		"Wallet:Read",
		"Wallet:List",
		"ServiceVehicleType:Read",
		"ServiceVehicleType:List",
		"ServiceProviderVehicle:Read",
		"ServiceProviderVehicle:List",
		"AppInstallation:Create",
		"AppInstallation:Update",
		"AppInstallation:Delete",
		"ServiceProvider:Create",
		"ServiceProvider:Update",
		"ServiceProvider:Delete",
		"User:Create",
		"User:Update",
		"User:Delete",
		"UserLocation:Create",
		"UserLocation:Update",
		"UserLocation:Delete",
		"ProviderLocation:Create",
		"ProviderLocation:Update",
		"ProviderLocation:Delete",
		"ServiceCompany:Create",
		"ServiceCompany:Update",
		"ServiceCompany:Delete",
		"ServiceProvider:Create",
		"ServiceProvider:Update",
		"ServiceProvider:Delete",
		"Service:Create",
		"Service:Update",
		"Service:Delete",
		"ServiceSubCategory:Create",
		"ServiceSubCategory:Update",
		"ServiceSubCategory:Delete",
		"ServiceType:Create",
		"ServiceType:Update",
		"ServiceType:Delete",
		"Coupon:Create",
		"Coupon:Update",
		"Coupon:Delete",
		"CancelReason:Create",
		"CancelReason:Update",
		"CancelReason:Delete",
		"Review:Create",
		"Review:Update",
		"Review:Delete",
		"PushNotification:Create",
		"PushNotification:Update",
		"PushNotification:Delete",
		"Page:Create",
		"Page:Update",
		"Page:Delete",
		"PackageType:Create",
		"PackageType:Update",
		"PackageType:Delete",
		"ServiceProviderVehicle:Create",
		"ServiceProviderVehicle:Update",
		"ServiceProviderVehicle:Delete",
		"ServiceVehicleType:Create",
		"ServiceVehicleType:Update",
		"ServiceVehicleType:Delete",
		"BookingFareEstimate:Create",
		"BookingFareEstimate:Update",
		"BookingFareEstimate:Delete",
		"AdvertisementBanner:Create",
		"AdvertisementBanner:Update",
		"AdvertisementBanner:Delete",
		"Store:Create",
		"Store:Update",
		"Store:Delete",
		"AppVersion:Create",
		"AppVersion:Update",
		"AppVersion:Delete",
		"DeliveryVehicleType:Create",
		"DeliveryVehicleType:Update",
		"DeliveryVehicleType:Delete",
		"StoreItemCategory:Create",
		"StoreItemCategory:Update",
		"StoreItemCategory:Delete",
		"StoreItem:Create",
		"StoreItem:Update",
		"StoreItem:Delete",
		"StoreItemType:Create",
		"StoreItemType:Update",
		"StoreItemType:Delete",
		"Order:Create",
		"Order:Update",
		"Order:Delete",
		"OrderStatusUtility:Create",
		"OrderStatusUtility:Update",
		"OrderStatusUtility:Delete",
		"DeliveryChargesUtility:Create",
		"DeliveryChargesUtility:Update",
		"DeliveryChargesUtility:Delete",
		"GeoFenceLocation:Create",
		"GeoFenceLocation:Update",
		"GeoFenceLocation:Delete",
		"GeoFenceRestrictedArea:Create",
		"GeoFenceRestrictedArea:Update",
		"GeoFenceRestrictedArea:Delete",
		"EmailTemplate:Create",
		"EmailTemplate:Update",
		"EmailTemplate:Delete",
		"SMSTemplate:Create",
		"SMSTemplate:Update",
		"SMSTemplate:Delete",
		"VehicleMake:Create",
		"VehicleMake:Update",
		"VehicleMake:Delete",
		"VehicleModel:Create",
		"VehicleModel:Update",
		"VehicleModel:Delete",
		"VisitLocation:Create",
		"VisitLocation:Update",
		"VisitLocation:Delete",
		"EnterpriseAccount:Create",
		"EnterpriseAccount:Update",
		"EnterpriseAccount:Delete",
		"RideProfileType:Create",
		"RideProfileType:Update",
		"RideProfileType:Delete",
		"BusinessTripReason:Create",
		"BusinessTripReason:Update",
		"BusinessTripReason:Delete",
		"Country:Create",
		"Country:Update",
		"Country:Delete",
		"State:Create",
		"State:Update",
		"State:Delete",
		"City:Create",
		"City:Update",
		"City:Delete",
		"File:Create",
		"File:Update",
		"File:Delete",
		"DeliveryCharge:Create",
		"DeliveryCharge:Update",
		"DeliveryCharge:Delete",
		"LocationWiseFare:Create",
		"LocationWiseFare:Update",
		"LocationWiseFare:Delete",
		"AirportSurcharge:Create",
		"AirportSurcharge:Update",
		"AirportSurcharge:Delete",
		"GeneralLabel:Create",
		"GeneralLabel:Update",
		"GeneralLabel:Delete",
		"FoodDeliveryLabel:Create",
		"FoodDeliveryLabel:Update",
		"FoodDeliveryLabel:Delete",
		"GroceryDeliveryLabel:Create",
		"GroceryDeliveryLabel:Update",
		"GroceryDeliveryLabel:Delete",
		"WineDeliveryLabel:Create",
		"WineDeliveryLabel:Update",
		"WineDeliveryLabel:Delete",
		"FAQ:Create",
		"FAQ:Update",
		"FAQ:Delete",
		"FAQCategory:Create",
		"FAQCategory:Update",
		"FAQCategory:Delete",
		"HelpDetail:Create",
		"HelpDetail:Update",
		"HelpDetail:Delete",
		"HelpCategory:Create",
		"HelpCategory:Update",
		"HelpCategory:Delete",
		"MarketSettings:Create",
		"MarketSettings:Update",
		"MarketSettings:Delete",
		"OAuthApplication:Create",
		"OAuthApplication:Update",
		"OAuthApplication:Delete",
		"AccessToken:Create",
		"AccessToken:Update",
		"AccessToken:Delete",
		"Webhook:Create",
		"Webhook:Update",
		"Webhook:Delete",
		"Currency:Create",
		"Currency:Update",
		"Currency:Delete",
		"RentalPackage:Create",
		"RentalPackage:Update",
		"RentalPackage:Delete",
		"StoreVehicleType:Create",
		"StoreVehicleType:Update",
		"StoreVehicleType:Delete",
		"RequiredDocument:Create",
		"RequiredDocument:Update",
		"RequiredDocument:Delete",
		"Document:Create",
		"Document:Update",
		"Document:Delete",
	}
	return serviceScopes
}
