package model

type Workspace struct {
	ID                 string
	Name               string
	Email              string
	SiloSpecifications []SiloSpecification
	SiloDefinitions    []SiloDefinition
}
