package behavior

import "fmt"

type Visitor interface {
	visitText(text TextElement)
	visitImage(image ImageElement)
}

type Element interface {
	accept(visitor Visitor)
}

type TextElement struct {
	text string
}

func (e *TextElement) accept(visitor Visitor) {
	visitor.visitText(*e)
}

type ImageElement struct {
	src string
}

func (e *ImageElement) accept(visitor Visitor) {
	visitor.visitImage(*e)
}

type SEOAnalyzer struct {
}

func (a *SEOAnalyzer) visitText(textElem TextElement) {
	fmt.Printf("analyzing SEO for text: %s", textElem.text)
}

func (a *SEOAnalyzer) visitImage(imageElem ImageElement) {
	fmt.Printf("analyzing SEO for image: %s", imageElem.src)
}

type SocialMediaAnalyzer struct {
}

func (a *SocialMediaAnalyzer) visitText(textElem TextElement) {
	fmt.Printf("analyzing Social media for text: %s", textElem.text)
}

func (a *SocialMediaAnalyzer) visitImage(imageElem ImageElement) {
	fmt.Printf("analyzing Social media for image: %s", imageElem.src)
}
func RunVisitor() {
	textElem := &TextElement{}
	imageElem := &ImageElement{}

	elements := []Element{textElem, imageElem}

	seoAnalyzer := &SEOAnalyzer{}
	socailMediaAnalyzer := &SocialMediaAnalyzer{}

	for _, elem := range elements {
		elem.accept(seoAnalyzer)
		elem.accept(socailMediaAnalyzer)
	}
}
