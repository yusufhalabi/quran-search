package types

type SearchResult struct {
	Verses []Verse `json:"verses"`
	Total  int     `json:"total"`
}

type Verse struct {
	ID          int    `json:"id"`
	SurahNumber int    `json:"surah_number"`
	AyahNumber  int    `json:"ayah_number"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}
