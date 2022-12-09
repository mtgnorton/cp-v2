package utility

var SensitiveInspector = &sensitiveInspector{
	replaceChar: '*',
	excludeChars: []rune{
		' ', '*', '&', '|', '"', '\'', '/', '<', '>', '?', '`', '~', '!', '@', '#', '$', '%', '^', '(', ')', '_', '+', '-', '=', '{', '}', '[', ']', ':', ';', '.', ',', '"',
	},
	root: &node{},
}

type sensitiveInspector struct {
	replaceChar  rune
	excludeChars []rune
	root         *node
}

func (d *sensitiveInspector) AddSensitiveWords(words []string) {
	for _, item := range words {
		d.root.addSensitiveWord(item)
	}
}

// MatchAndReplace 查找替换发现的敏感词
func (d *sensitiveInspector) MatchAndReplace(text string) (sensitiveWords []string, replaceText string) {
	if d.root == nil {
		return nil, text
	}

	textChars := []rune(text)
	textCharsCopy := make([]rune, len(textChars))

	copy(textCharsCopy, textChars)

	length := len(textChars)

	for i := 0; i < length; i++ {

		if d.checkExclude(textChars[i]) {
			continue
		}
		//root本身是没有key的，root的下面一个节点，才算是第一个；
		currentNode := d.root.findCharacter(textChars[i])
		if currentNode == nil {
			continue
		}

		j := i + 1
		for ; j <= length && currentNode != nil; j++ {

			if currentNode.End {
				sensitiveWords = append(sensitiveWords, string(textChars[i:j]))
				d.replaceRune(textCharsCopy, i, j)
			}
			if j >= length {
				break
			}
			if d.checkExclude(textChars[j]) {
				continue
			}
			currentNode = currentNode.findCharacter(textChars[j])
		}
	}
	return sensitiveWords, string(textCharsCopy)
}

func (d *sensitiveInspector) checkExclude(u rune) bool {
	if len(d.excludeChars) == 0 {
		return false
	}
	var exist bool
	for i, l := 0, len(d.excludeChars); i < l; i++ {
		if u == d.excludeChars[i] {
			exist = true
			break
		}
	}
	return exist
}

func (d *sensitiveInspector) replaceRune(characters []rune, begin int, end int) {
	for i := begin; i < end; i++ {
		characters[i] = d.replaceChar
	}
}

type node struct {
	End  bool
	Next map[rune]*node
}

func (n *node) addCharacter(c rune) *node {
	if n.Next == nil {
		n.Next = make(map[rune]*node)
	}
	//如果已经存在了，就不再往里面添加了；
	if next, ok := n.Next[c]; ok {
		return next
	} else {
		n.Next[c] = &node{
			End:  false,
			Next: nil,
		}
		return n.Next[c]
	}
}

func (n *node) findCharacter(c rune) *node {
	if n.Next == nil {
		return nil
	}

	if _, ok := n.Next[c]; ok {
		return n.Next[c]
	}
	return nil
}

func (n *node) addSensitiveWord(word string) {
	node := n
	characters := []rune(word)
	for index, _ := range characters {
		node = node.addCharacter(characters[index])
	}
	node.End = true
}
