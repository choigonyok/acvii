package parser

import (
	"os"
	"strings"

	"github.com/choigonyok/acvii/pkg/watcher"
	"github.com/fsnotify/fsnotify"
)

type Parser struct {
	watcher                 *watcher.Watcher
	authorizationPolicyPool []string
}

// New creates new parser that parse changed AuthorizationPolicy manifest file
func New(w *watcher.Watcher) *Parser {
	p := &Parser{watcher: w}
	for _, v := range w.GetWatchList() {
		p.parseManifest(v)
	}
	return p
}

// Parse parses AuthorizationPolicy manifest, resets and restores AuthorizationPolicy pool
func (p *Parser) Parse(e fsnotify.Event) {
	if e.Has(fsnotify.Write) {
		p.resetAuthorizationPolicyPool()
		for _, v := range p.watcher.GetWatchList() {
			p.parseManifest(v)
		}
	}
}

// ParseManifest split manifest's each AuthorizationPolicy resource and append them to AuthorizationPolicy pool
func (p *Parser) parseManifest(filepath string) {
	b, _ := os.ReadFile(filepath)
	fileContent := string(b)
	fileContent, _ = strings.CutPrefix(fileContent, "---")
	fileContent, _ = strings.CutSuffix(fileContent, "---")
	p.authorizationPolicyPool = append(p.authorizationPolicyPool, strings.Split(fileContent, "---")...)

	// viper.AddConfigPath(m.Path)
	// viper.SetConfigName(m.Name)
	// viper.SetConfigType(m.Extension)
	// viper.ReadInConfig()
	// for _, v := range viper.AllKeys() {
	// 	fmt.Printf("%v : %v\n", viper.AllKeys(), viper.Get(v))
	// }
}

// TEST:
// GetManifestFromPool returns specific AuthorizationPolicy manifest in pool by index
func (p *Parser) GetManifestFromPool(index int) string {
	return p.authorizationPolicyPool[index]
}

// resetAuthorizationPolicyPool resets current AuthorizationPolicy pool
func (p *Parser) resetAuthorizationPolicyPool() {
	p.authorizationPolicyPool = nil
}
