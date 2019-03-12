# notihub/rssgen

rssgen is a Go application that helps you generate RSS on your website using simple conditional expressions.

## Goals
- Provides a simple interface to generate RSS on website

## Usage
```go
package main

import (
	"fmt"
	"github.com/notihub/rssgen"
)

func main() {
	rule := `<a href="{%:link}">
                      <span class="text-normal">
                        {%:name} / 
                      </span>
                      {%:title}
                    </a>
                  {*}
                  <p class="col-9 d-inline-block text-gray m-0 pr-4">
                            {%:description}
                  </p>`
	rss, _ := rssgen.GenerateRSS(rssgen.Generator{Url: "https://github.com/trending", Rule: rule})
	fmt.Println(rss)
}
```

Outputs:

```xml
<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title></title>
  <id>https://github.com/trending</id>
  <updated>2019-03-12T23:46:18+09:00</updated>
  <link href="https://github.com/trending"></link>
  <entry>
    <title>takeover.sh</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/#start-of-content</id>
    <link href="https://github.com/#start-of-content" rel="alternate"></link>
    <summary type="html">Wipe and reinstall a running Linux system via SSH, without rebooting. You know you want to.</summary>
    <author>
      <name>marcan</name>
    </author>
  </entry>
  <entry>
    <title>code-server</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/codercom/code-server</id>
    <link href="https://github.com/codercom/code-server" rel="alternate"></link>
    <summary type="html">Run VS Code on a remote server.</summary>
    <author>
      <name>codercom</name>
    </author>
  </entry>
  <entry>
    <title>calculator</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/Microsoft/calculator</id>
    <link href="https://github.com/Microsoft/calculator" rel="alternate"></link>
    <summary type="html">Windows Calculator: A simple yet powerful calculator that ships with Windows</summary>
    <author>
      <name>Microsoft</name>
    </author>
  </entry>
  <entry>
    <title>ghidra</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/NationalSecurityAgency/ghidra</id>
    <link href="https://github.com/NationalSecurityAgency/ghidra" rel="alternate"></link>
    <summary type="html">Ghidra is a software reverse engineering (SRE) framework</summary>
    <author>
      <name>NationalSecurityAgency</name>
    </author>
  </entry>
  <entry>
    <title>spring-boot-examples</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/ityouknow/spring-boot-examples</id>
    <link href="https://github.com/ityouknow/spring-boot-examples" rel="alternate"></link>
    <summary type="html">about learning Spring Boot via examples. Spring Boot 教程、技术栈示例代码，快速简单上手教程。</summary>
    <author>
      <name>ityouknow</name>
    </author>
  </entry>
  <entry>
    <title>server</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/gotify/server</id>
    <link href="https://github.com/gotify/server" rel="alternate"></link>
    <summary type="html">A simple server for sending and receiving messages in real-time per web socket. (Includes a sleek web-ui)</summary>
    <author>
      <name>gotify</name>
    </author>
  </entry>
  <entry>
    <title>lets-get-arrested</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/hamukazu/lets-get-arrested</id>
    <link href="https://github.com/hamukazu/lets-get-arrested" rel="alternate"></link>
    <summary type="html">This project is intended to protest against the police in Japan</summary>
    <author>
      <name>hamukazu</name>
    </author>
  </entry>
  <entry>
    <title>ArchiveBox</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/pirate/ArchiveBox</id>
    <link href="https://github.com/pirate/ArchiveBox" rel="alternate"></link>
    <summary type="html">Translation management platform for teams</summary>
    <author>
      <name>pirate</name>
    </author>
  </entry>
  <entry>
    <title>react-three-fiber</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/drcmda/react-three-fiber</id>
    <link href="https://github.com/drcmda/react-three-fiber" rel="alternate"></link>
    <summary type="html">A flowchart that explains the new lifecycle of a Hooks component.</summary>
    <author>
      <name>drcmda</name>
    </author>
  </entry>
  <entry>
    <title>curv</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/curv3d/curv</id>
    <link href="https://github.com/curv3d/curv" rel="alternate"></link>
    <summary type="html">a language for making art using mathematics</summary>
    <author>
      <name>curv3d</name>
    </author>
  </entry>
  <entry>
    <title>Go42</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/ffhelicopter/Go42</id>
    <link href="https://github.com/ffhelicopter/Go42" rel="alternate"></link>
    <summary type="html">写《Go语言四十二章经》，纯粹是因为开发过程中碰到过的一些问题，踩到过的一些坑，感觉在Go语言学习使用过程中，有必要深刻理解这门语言的核心思维、清晰掌握语言的细节规范以及反复琢磨标准包代码设计模式，于是才有了这本书。</summary>
    <author>
      <name>ffhelicopter</name>
    </author>
  </entry>
  <entry>
    <title>JavaGuide</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/Snailclimb/JavaGuide</id>
    <link href="https://github.com/Snailclimb/JavaGuide" rel="alternate"></link>
    <summary type="html">【Java学习+面试指南】 一份涵盖大部分Java程序员所需要掌握的核心知识。</summary>
    <author>
      <name>Snailclimb</name>
    </author>
  </entry>
  <entry>
    <title>CS-Notes</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/CyC2018/CS-Notes</id>
    <link href="https://github.com/CyC2018/CS-Notes" rel="alternate"></link>
    <summary type="html">React components for efficiently rendering large lists and tabular data</summary>
    <author>
      <name>CyC2018</name>
    </author>
  </entry>
  <entry>
    <title>Micro8</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/Micropoor/Micro8</id>
    <link href="https://github.com/Micropoor/Micro8" rel="alternate"></link>
    <summary type="html">Gitbook</summary>
    <author>
      <name>Micropoor</name>
    </author>
  </entry>
  <entry>
    <title>personal-website</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/github/personal-website</id>
    <link href="https://github.com/github/personal-website" rel="alternate"></link>
    <summary type="html">Code that&#39;ll help you kickstart a personal website that showcases your work as a software developer.</summary>
    <author>
      <name>github</name>
    </author>
  </entry>
  <entry>
    <title>pytorch_geometric</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/rusty1s/pytorch_geometric</id>
    <link href="https://github.com/rusty1s/pytorch_geometric" rel="alternate"></link>
    <summary type="html">Geometric Deep Learning Extension Library for PyTorch</summary>
    <author>
      <name>rusty1s</name>
    </author>
  </entry>
  <entry>
    <title>cmake-examples</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/pr0g/cmake-examples</id>
    <link href="https://github.com/pr0g/cmake-examples" rel="alternate"></link>
    <summary type="html">A collection of as simple as possible, modern CMake projects</summary>
    <author>
      <name>pr0g</name>
    </author>
  </entry>
  <entry>
    <title>slim</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/openacid/slim</id>
    <link href="https://github.com/openacid/slim" rel="alternate"></link>
    <summary type="html">Unbelievably space efficient data structures in Golang.</summary>
    <author>
      <name>openacid</name>
    </author>
  </entry>
  <entry>
    <title>freeCodeCamp</title>
    <updated>2019-03-12T23:46:18+09:00</updated>
    <id>https://github.com/freeCodeCamp/freeCodeCamp</id>
    <link href="https://github.com/freeCodeCamp/freeCodeCamp" rel="alternate"></link>
    <summary type="html">The &#xA;                    &lt;a href=&#34;https://www.freeCodeCamp.org&#34; rel=&#34;nofollow&#34;&gt;&#xA;                      https://www.freeCodeCamp.org&#xA;                    &lt;/a&gt;&#xA;                     open source codebase and curriculum. Learn to code for free together with millions of people.</summary>
    <author>
      <name>freeCodeCamp</name>
    </author>
  </entry>
</feed>
```