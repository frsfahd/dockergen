package catalog

// supported programming languages
type Language string

const (
	Node   Language = "node"
	Python Language = "python"
)

// supported frameworks
type NodeFramework string
type PythonFramework string

const (
	NodeDefault NodeFramework = "node"
	NextJS      NodeFramework = "next"
	Vite        NodeFramework = "vite"
	ExpressJS   NodeFramework = "express"
	Hapi        NodeFramework = "hapi"
	Fastify     NodeFramework = "fastify"
)

const (
	PythonDefault PythonFramework = "python"
	Django        PythonFramework = "django"
	FastAPI       PythonFramework = "fastapi"
	Flask         PythonFramework = "flask"
)
