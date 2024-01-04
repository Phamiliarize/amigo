package dto

type ThemeMetadata struct {
	CommunityName string
	Description   string
	DefaultTheme  string
}

type Theme struct {
	Name         string
	Publish      bool
	Path         string
	Dir          string
	BaseThemeDir string
	ReadingMode  bool
}

type CachedThemeMetadata struct {
	CachedAt      int64
	Settings      GeneralSetting
	ThemeMetadata ThemeMetadata
}
