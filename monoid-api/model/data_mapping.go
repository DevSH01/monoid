package model

type SiloSpecification struct {
	ID              string
	Name            string
	LogoURL         *string
	WorkspaceID     string
	Workspace       Workspace `gorm:"constraint:OnDelete:CASCADE;"`
	DockerImage     *string
	Schema          *string
	SiloDefinitions []SiloDefinition
}

type SiloDefinition struct {
	ID                  string
	WorkspaceID         string
	Workspace           Workspace `gorm:"constraint:OnDelete:CASCADE;"`
	Description         *string
	SiloSpecificationID string
	SiloSpecification   SiloSpecification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DataSources         []DataSource
	Subjects            []Subject `gorm:"many2many:silo_definition_subjects;"`
}

type DataSource struct {
	ID               string
	SiloDefinitionID string
	SiloDefinition   SiloDefinition `gorm:"constraint:OnDelete:CASCADE;"`
	Properties       []*Property
	Description      *string
	Schema           string
}

type Property struct {
	ID           string      `json:"id"`
	Categories   []*Category `json:"categories"`
	DataSourceID string      `json:"dataSourceID"`
	DataSource   DataSource  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Purposes     []*Purpose  `json:"purposes"`
}

type Subject struct {
	ID              string
	Name            string
	WorkspaceID     string
	Workspace       Workspace `gorm:"constraint:OnDelete:CASCADE;"`
	SiloDefinitions []SiloDefinition
}

type Category struct {
	ID          string
	Name        string
	WorkspaceID string
	Workspace   Workspace `gorm:"constraint:OnDelete:CASCADE;"`
	Properties  []Property
}

type Purpose struct {
	ID          string
	Name        string
	WorkspaceID string
	Workspace   Workspace `gorm:"constraint:OnDelete:CASCADE;"`
	Properties  []Property
}
