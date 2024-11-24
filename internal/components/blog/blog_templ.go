// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package blog

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/gmeylan/go-website/internal/components/layout"

func Blog() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen bg-gray-50\"><main class=\"max-w-7xl mx-auto py-16 px-4 sm:px-6 lg:px-8\"><div class=\"space-y-16\"><section><div class=\"text-center\"><h1 class=\"text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl\">Blog</h1><p class=\"mt-3 max-w-2xl mx-auto text-xl text-gray-500\">Réflexions sur le développement web, DevOps et les bonnes pratiques</p></div></section><section class=\"bg-white shadow-lg rounded-2xl overflow-hidden\"><div class=\"p-6 sm:p-8\"><div class=\"flex flex-wrap gap-3\"><button class=\"px-4 py-2 rounded-md bg-indigo-600 text-white hover:bg-indigo-700\">Tous les articles</button> <button class=\"px-4 py-2 rounded-md bg-white border border-gray-300 text-gray-700 hover:bg-gray-50\">DevOps</button> <button class=\"px-4 py-2 rounded-md bg-white border border-gray-300 text-gray-700 hover:bg-gray-50\">Go</button> <button class=\"px-4 py-2 rounded-md bg-white border border-gray-300 text-gray-700 hover:bg-gray-50\">Web</button></div></div></section><section class=\"space-y-8\"><h2 class=\"text-2xl font-bold text-gray-900\">Articles récents</h2><div class=\"grid gap-8 md:grid-cols-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderFeaturedPost().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderFeaturedPost().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></section><section class=\"space-y-8\"><h2 class=\"text-2xl font-bold text-gray-900\">Tous les articles</h2><div class=\"grid gap-8 sm:grid-cols-2 lg:grid-cols-3\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = renderPostCard().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></section></div></main></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Page("Blog index").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func renderFeaturedPost() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article class=\"bg-white shadow-lg rounded-2xl overflow-hidden hover:shadow-xl transition-shadow\"><div class=\"aspect-w-16 aspect-h-9\"><img src=\"https://placehold.co/800x400/png\" alt=\"Article image\" class=\"object-cover w-full h-full\"></div><div class=\"p-6\"><div class=\"flex gap-2 mb-4\"><span class=\"px-3 py-1 text-sm font-medium text-indigo-600 bg-indigo-100 rounded-md\">Go</span> <span class=\"px-3 py-1 text-sm font-medium text-indigo-600 bg-indigo-100 rounded-md\">Web</span></div><h3 class=\"text-xl font-bold text-gray-900 mb-2\"><a href=\"/blog/slug-de-ouf\" class=\"hover:text-indigo-600\">Introduction à HTMX avec Go</a></h3><p class=\"text-gray-600 mb-4\">Découvrez comment HTMX peut simplifier le développement web moderne avec Go et créer des applications interactives sans JavaScript complexe.</p><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"ml-3\"><p class=\"text-sm font-medium text-gray-900\">John Doe</p><p class=\"text-sm text-gray-500\">15 Mars 2024</p></div></div></div></div></article>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func renderPostCard() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article class=\"bg-white shadow-lg rounded-2xl overflow-hidden hover:shadow-xl transition-shadow\"><div class=\"aspect-w-16 aspect-h-9\"><img src=\"https://placehold.co/400x225/png\" alt=\"Article image\" class=\"object-cover w-full h-full\"></div><div class=\"p-6\"><div class=\"flex gap-2 mb-3\"><span class=\"px-3 py-1 text-xs font-medium text-indigo-600 bg-indigo-100 rounded-md\">Docker</span></div><h3 class=\"text-lg font-bold text-gray-900 mb-2\"><a href=\"#\" class=\"hover:text-indigo-600\">Optimiser ses builds Docker</a></h3><p class=\"text-gray-600 text-sm mb-4\">Apprenez les meilleures pratiques pour optimiser vos builds Docker et améliorer vos workflows de développement.</p><div class=\"flex items-center justify-between text-sm\"><div class=\"flex items-center\"><span class=\"ml-2 text-gray-700\">Jane Smith</span></div><span class=\"text-gray-500\">12 Mars 2024</span></div></div></article>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
