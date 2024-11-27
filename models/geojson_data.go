package models

type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"` // Adjust according to your specific needs (nested arrays for polygons)
}

type Properties struct {
	ShapeName  string `json:"shapeName"`
	ShapeISO   string `json:"shapeISO"`
	ShapeID    string `json:"shapeID"`
	ShapeGroup string `json:"shapeGroup"`
	ShapeType  string `json:"shapeType"`
}
