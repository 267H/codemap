package config

type Config struct {
	MaxFileSizeBytes      int
	OutputFileName        string
	WarningFilesThreshold int
	WarningDirsThreshold  int
	TokensPerChar         float64
	ExcludeDirs           map[string]bool
	ExcludeExtensions     map[string]bool
	FileExtensionMap      map[string]FileExtensionInfo
}

type FileExtensionInfo struct {
	IsCode   bool
	Language string
}

func NewDefaultConfig() *Config {
	cfg := &Config{
		MaxFileSizeBytes:      1024 * 1024,
		OutputFileName:        "codebase_map.txt",
		WarningFilesThreshold: 1000,
		WarningDirsThreshold:  100,
		TokensPerChar:         0.25,
		ExcludeDirs: map[string]bool{
			".git":         true,
			"node_modules": true,
			"vendor":       true,
			"dist":         true,
			"build":        true,
		},
		ExcludeExtensions: map[string]bool{
			".exe":   true,
			".dll":   true,
			".so":    true,
			".dylib": true,
			".bin":   true,
			".obj":   true,
			".o":     true,
			".a":     true,
			".lib":   true,
			".pyc":   true,
			".pyo":   true,
			".class": true,
		},
		FileExtensionMap: make(map[string]FileExtensionInfo),
	}

	initializeExtensionMap(cfg)
	return cfg
}

func initializeExtensionMap(cfg *Config) {
	extMap := map[string]FileExtensionInfo{
		".go":     {true, "go"},
		".js":     {true, "javascript"},
		".jsx":    {true, "jsx"},
		".ts":     {true, "typescript"},
		".tsx":    {true, "tsx"},
		".py":     {true, "python"},
		".java":   {true, "java"},
		".c":      {true, "c"},
		".cpp":    {true, "cpp"},
		".cc":     {true, "cpp"},
		".h":      {true, "c"},
		".hpp":    {true, "cpp"},
		".cs":     {true, "csharp"},
		".rb":     {true, "ruby"},
		".php":    {true, "php"},
		".html":   {true, "html"},
		".css":    {true, "css"},
		".scss":   {true, "scss"},
		".sql":    {true, "sql"},
		".swift":  {true, "swift"},
		".kt":     {true, "kotlin"},
		".rs":     {true, "rust"},
		".sh":     {true, "bash"},
		".bash":   {true, "bash"},
		".pl":     {true, "perl"},
		".json":   {true, "json"},
		".yaml":   {true, "yaml"},
		".yml":    {true, "yaml"},
		".xml":    {true, "xml"},
		".md":     {true, "markdown"},
		".txt":    {true, "plaintext"},
		".toml":   {true, "toml"},
		".ini":    {true, "ini"},
		".cfg":    {true, "ini"},
		".proto":  {true, "protobuf"},
		".dart":   {true, "dart"},
		".lua":    {true, "lua"},
		".ex":     {true, "elixir"},
		".exs":    {true, "elixir"},
		".erl":    {true, "erlang"},
		".hs":     {true, "haskell"},
		".ml":     {true, "ocaml"},
		".scala":  {true, "scala"},
		".clj":    {true, "clojure"},
		".fs":     {true, "fsharp"},
		".r":      {true, "r"},
		".groovy": {true, "groovy"},
		".jl":     {true, "julia"},
		".d":      {true, "d"},
		".zig":    {true, "zig"},
		".odin":   {true, "odin"},
		".nim":    {true, "nim"},
		".v":      {true, "v"},
		".asm":    {true, "assembly"},
		".s":      {true, "assembly"},
		".elm":    {true, "elm"},
		".f":      {true, "fortran"},
		".f90":    {true, "fortran"},
		".f95":    {true, "fortran"},
		".mat":    {true, "matlab"},
		".m":      {true, "objective-c"},
		".mm":     {true, "objective-cpp"},
		".pas":    {true, "pascal"},
		".pp":     {true, "pascal"},
		".cob":    {true, "cobol"},
		".lisp":   {true, "lisp"},
		".cl":     {true, "lisp"},
		".bas":    {true, "basic"},
	}

	cfg.FileExtensionMap = extMap
}
