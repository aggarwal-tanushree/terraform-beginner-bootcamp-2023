/* package main: Declares the package name. 
The main package is special in Go, it's where the execution of the program starts.
tells the Go compiler that the package should compile as an executable program instead of a shared library package main*/
package main

// fmt is short format, it contains functions for formatted I/O.
import (
	"log"  // used for logging in validation func
	"fmt"  // used in main to print std output
	"context"// for providerConfigure func
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag" // for providerConfigure func
	"github.com/google/uuid"  // used by validation func for validating the UUID format
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema" // for ResourcesMap and DataSourcesMap in Provider func
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"   // being used in main func by plugin.Serve(&plugin.ServeOpts
)
// func main(): Defines the main function, the entry point of the app. 
// When you run the program, it starts executing from this function.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,  //calling provider function(or interface in goLang)
	})
	// Format.PrintLine
	// Prints to standard output
	fmt.Println("Hello, world!")
}

type Config struct {
	Endpoint string
	Token string
	UserUuid string
}

// in golang, a titlecase function will get exported.
func Provider() *schema.Provider {
	var p *schema.Provider // defined a variable p which is going to reference the provider
	p = &schema.Provider{
		ResourcesMap:  map[string]*schema.Resource{ // maps to actual resources
			"terratowns_home": Resource(),
		},
		DataSourcesMap:  map[string]*schema.Resource{  // maps to fields of a resource

		},
		Schema: map[string]*schema.Schema{ // json format in which our HTTP request will send the data
			"endpoint": {
				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for the external service",
			},
			"token": {
				Type: schema.TypeString,
				Sensitive: true, // make the token as sensitive to hide it the logs
				Required: true,
				Description: "Bearer token for authorization",
			},
			"user_uuid": {
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
				ValidateFunc: validateUUID,  // validation func to ensure UUID is in the expected format
			},
		},
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

// validation func to ensure UUID is in the expected format
func validateUUID(v interface{}, k string) (ws []string, errors []error) {  //array of string and array of errors
	log.Print("validateUUID:start")  //logging 
	value := v.(string)
	if _, err := uuid.Parse(value); err != nil {
		errors = append(errors, fmt.Errorf("invalid UUID format"))
	}
	log.Print("validateUUID:end")
	return
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {  //definition of func providerConfigure, takes argument pointer p
	// p is pomter to Provider object of TF schema package
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics ) {
		log.Print("providerConfigure:start")
		config := Config{ //variable named config. Config is a struct defined in our code
			Endpoint: d.Get("endpoint").(string), // Define endpint
			Token: d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfigure:end")
		return &config, nil
	}
}

// definition of Resource func
func Resource() *schema.Resource {
	log.Print("Resource:start")
	resource := &schema.Resource{
		// defining four functions (aka interfaces in Go) for the four actions
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
	}
	log.Print("Resource:end")
	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics // related to error handling
	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}	

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}