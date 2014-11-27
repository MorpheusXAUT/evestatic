// mysqllookup
package main

import (
	"log"

	"github.com/morpheusxaut/evestatic"
)

func main() {
	eve, err := evestatic.NewEveStatic(evestatic.DatasourceTypeMySql, "evetest:evetest@tcp(network.morpheusxaut.net:3316)/eveonline")
	if err != nil {
		log.Fatalf("Failed to create new MySQL-based EveStatic: [%v]", err)
		return
	}

	log.Println("Successfully created new MySQL-based EveStatic")

	id, err := eve.GetTypeIDFromName("Prostitute")
	if err != nil || id < 0 {
		log.Fatalf("Failed to retrieve typeID for 'Prostitute': [%v]", err)
	}

	log.Printf("Successfully retrieved typeID for 'Prostitute': %d\n", id)

	name, err := eve.GetTypeNameFromID(id)
	if err != nil || name == "" {
		log.Fatalf("Failed to retrieve typeName for ID %d: [%v]", id, err)
	}

	log.Printf("Successfully retrieved typeName for ID %d: %q\n", id, name)

	eve.Close()
}
