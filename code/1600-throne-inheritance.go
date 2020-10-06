package main

type ThroneInheritance struct {
	king     string
	children map[string][]string
	dead     map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{king: kingName, children: make(map[string][]string), dead: make(map[string]bool)}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.children[parentName] = append(this.children[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	res := []string{}
	this.dfs(this.king, &res)
	return res
}

func (this *ThroneInheritance) dfs(name string, res *[]string) {
	if !this.dead[name] {
		*res = append(*res, name)
	}
	for _, c := range this.children[name] {
		this.dfs(c, res)
	}
}
