package document

func addParams(pNode *Node) {
	qNode := newNode(5, 7, MapNode)
	pNode.addChild(qNode)
	envNode := newNode(8, 21, MapNode)
	pNode.addChild(envNode)
	qNode.addChild(newNode(6, 6, ScalarNode))
	qNode.addChild(newNode(7, 7, ScalarNode))
	envNode.addChild(newNode(9, 9, ScalarNode))
	envNode.addChild(newNode(10, 10, ScalarNode))
	envNode.addChild(newNode(11, 11, ScalarNode))
	allNode := newNode(12, 20, ListNode)
	envNode.addChild(allNode)
	allNode.addChild(newNode(13, 13, ScalarNode))
	allNode.addChild(newNode(14, 14, ScalarNode))
	allNode.addChild(newNode(15, 15, ScalarNode))
	allNode.addChild(newNode(16, 16, ScalarNode))
	allNode.addChild(newNode(17, 17, ScalarNode))
	allNode.addChild(newNode(18, 18, ScalarNode))
	allNode.addChild(newNode(19, 19, ScalarNode))
	allNode.addChild(newNode(20, 20, ScalarNode))
	envNode.addChild(newNode(21, 21, ScalarNode))
}

func addResources(rNode *Node) {
	qNode := newNode(24, 30, MapNode)
	rNode.addChild(qNode)
	qNode.addChild(newNode(25, 25, ScalarNode))
	propNode := newNode(26, 30, MapNode)
	qNode.addChild(propNode)
	propNode.addChild(newNode(27, 27, ScalarNode))
	propNode.addChild(newNode(28, 28, ScalarNode))
	visNode := newNode(29, 30, MapNode)
	propNode.addChild(visNode)
	visNode.addChild(newNode(30, 30, ScalarNode))
}

//hardcoded sps-V1.2.yaml
func hardCodedTree() *Tree {
	root := newNode(1, 35, RootNode)
	root.Children = append(
		root.Children,
		newNode(1, 1, ScalarNode),
	)
	root.Children = append(
		root.Children,
		newNode(2, 2, ScalarNode),
	)
	paramNode := newNode(4, 21, MapNode)
	root.addChild(paramNode)
	resNode := newNode(23, 30, MapNode)
	root.addChild(resNode)
	addParams(paramNode)
	addResources(resNode)
	qArnNode := newNode(33, 35, MapNode)
	outNode := newNode(32, 35, MapNode)
	root.addChild(outNode)
	outNode.addChild(qArnNode)
	qArnNode.addChild(newNode(34, 34, ScalarNode))
	qArnNode.addChild(newNode(35, 35, ScalarNode))
	return (*Tree)(root)
}
