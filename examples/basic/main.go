package main

import (
	"fmt"
	"log"
	"os"

	yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
)

func main() {
	// Get OAuth token from environment variable
	oauthToken := os.Getenv("YANDEX_CLOUD_OAUTH_TOKEN")
	if oauthToken == "" {
		log.Fatal("YANDEX_CLOUD_OAUTH_TOKEN environment variable is required")
	}

	// Initialize client
	client, err := yandexcloud.NewClient(oauthToken, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("✅ Client initialized successfully")

	// List organizations
	fmt.Println("\n📋 Listing organizations...")
	organizations, err := client.Organizations().List(nil, nil)
	if err != nil {
		log.Fatalf("Failed to list organizations: %v", err)
	}
	fmt.Printf("Organizations: %+v\n", organizations)

	// Get organization ID from environment or use first from list
	organizationID := os.Getenv("YANDEX_CLOUD_ORGANIZATION_ID")
	if organizationID == "" && organizations["organizations"] != nil {
		orgs := organizations["organizations"].([]interface{})
		if len(orgs) > 0 {
			org := orgs[0].(map[string]interface{})
			organizationID = org["id"].(string)
		}
	}

	if organizationID != "" {
		// List clouds in organization
		fmt.Printf("\n☁️  Listing clouds in organization %s...\n", organizationID)
		clouds, err := client.Clouds().List(&organizationID, nil, nil)
		if err != nil {
			log.Fatalf("Failed to list clouds: %v", err)
		}
		fmt.Printf("Clouds: %+v\n", clouds)

		// Get cloud ID from environment or use first from list
		cloudID := os.Getenv("YANDEX_CLOUD_CLOUD_ID")
		if cloudID == "" && clouds["clouds"] != nil {
			cloudsList := clouds["clouds"].([]interface{})
			if len(cloudsList) > 0 {
				cloud := cloudsList[0].(map[string]interface{})
				cloudID = cloud["id"].(string)
			}
		}

		if cloudID != "" {
			// List folders in cloud
			fmt.Printf("\n📁 Listing folders in cloud %s...\n", cloudID)
			folders, err := client.Folders().List(cloudID, nil, nil)
			if err != nil {
				log.Fatalf("Failed to list folders: %v", err)
			}
			fmt.Printf("Folders: %+v\n", folders)
		}
	}

	fmt.Println("\n✨ Example completed successfully!")
}
