package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Category",
            Router: "/category/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: "/category/addcategory",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CategoryTitle",
            Router: "/categorytitle",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetWxProjectConfig",
            Router: "/getwxprojectconfig",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Jsoneditor",
            Router: "/jsoneditor",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
        beego.ControllerComments{
            Method: "PutWxProjectConfig",
            Router: "/putwxprojectconfig",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
        beego.ControllerComments{
            Method: "ErrLog",
            Router: "/errlog",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
        beego.ControllerComments{
            Method: "InfoLog",
            Router: "/infolog",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "AddAnsysArticles",
            Router: "/addansysarticle/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysArticle",
            Router: "/ansysarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "AnsysCal",
            Router: "/ansyscal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "DeleteAnsys",
            Router: "/deleteansys/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "EditorAnsysArticle",
            Router: "/editoransysarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxAnsysHistory",
            Router: "/getWxansyshistory/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsys",
            Router: "/getansys/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "AnsysList",
            Router: "/getansys/ansyslist/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsyss",
            Router: "/getansys/getansyss/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxAnsyss",
            Router: "/getansys/getwxansyss/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysCalInput",
            Router: "/getansyscalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetansysCalOutput",
            Router: "/getansyscaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysHistory",
            Router: "/getansyshistory/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysHistoryCal",
            Router: "/getansyshistorycal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysHistoryInput",
            Router: "/getansyshistoryinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysHistoryOutput",
            Router: "/getansyshistoryoutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetAnsysPdf",
            Router: "/getansyspdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetHistoryAnsys",
            Router: "/gethistoryansys/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxansysCalInput",
            Router: "/getwxansyscalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxAnsysCalOutput",
            Router: "/getwxansyscaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxAnsysClass",
            Router: "/getwxansysclass/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "PostAnsys",
            Router: "/postansys",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "PostWxansys",
            Router: "/postwxansys",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "PutAnsysApdl",
            Router: "/putansysapdl",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "PutansysCalInput",
            Router: "/putansyscalinput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "PutansysCalOutput",
            Router: "/putansyscaloutput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "UpdateAnsysArticle",
            Router: "/updateansysarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "UploadAnsys",
            Router: "/uploadansys/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "UploadAnsysApdl",
            Router: "/uploadansysapdl/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AnsysController"],
        beego.ControllerComments{
            Method: "GetWxAnsysArticle",
            Router: "/wxansysarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddWxArticle",
            Router: "/addwxarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddWxArticleFlow",
            Router: "/addwxarticleflow/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddWxArticles",
            Router: "/addwxarticles/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddWxEditorArticle",
            Router: "/addwxeditorarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "DeleteWxArticle",
            Router: "/deletewxarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetListArticles",
            Router: "/getlistarticles",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetUserArticle",
            Router: "/getuserarticle",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxArticle",
            Router: "/getwxarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxArticleFlow",
            Router: "/getwxarticleflow/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxArticles",
            Router: "/getwxarticles",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxArticless",
            Router: "/getwxarticless/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxArticleType",
            Router: "/getwxarticletype",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetWxUserArticles",
            Router: "/getwxuserarticles/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "SearchWxArticles",
            Router: "/searchwxarticles/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "UpdateWxEditorArticle",
            Router: "/updatewxeditorarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "AddWxAttachment",
            Router: "/addwxattachment",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetExcelPdf",
            Router: "/getexcelpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetMathPdf",
            Router: "/getmathpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetWxExcelPdf",
            Router: "/getwxexcelpdf",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetWxExcelTempPdf",
            Router: "/getwxexceltemppdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetWxMathPdf",
            Router: "/getwxmathpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetWxPdf",
            Router: "/getwxpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetWxTempPdf",
            Router: "/getwxtemppdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "GetPdfs",
            Router: "/project/product/pdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AttachController"],
        beego.ControllerComments{
            Method: "WxPdf",
            Router: "/wxpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"],
        beego.ControllerComments{
            Method: "Bbs",
            Router: "/bbs",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"],
        beego.ControllerComments{
            Method: "BbsGetBbs",
            Router: "/bbsgetbbs",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BbsController"],
        beego.ControllerComments{
            Method: "GetBbs",
            Router: "/getbbs",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "AddBusiness",
            Router: "/addbusiness/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessCheck",
            Router: "/businesscheck",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheck",
            Router: "/businessmonthcheck/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheck2",
            Router: "/businessmonthcheck2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheck3",
            Router: "/businessmonthcheck3/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheck4",
            Router: "/businessmonthcheck4/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheck5",
            Router: "/businessmonthcheck5/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheckSum",
            Router: "/businessmonthchecksum",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheckSum2",
            Router: "/businessmonthchecksum2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheckSum3",
            Router: "/businessmonthchecksum3",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheckSum4",
            Router: "/businessmonthchecksum4",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "BusinessMonthCheckSum5",
            Router: "/businessmonthchecksum5",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "GetBysiness",
            Router: "/getbusiness/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "GetBysinessById",
            Router: "/getbusinessbyid/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "GetBusinessCheck",
            Router: "/getbusinesscheck",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:BusinessController"],
        beego.ControllerComments{
            Method: "UpdateBusiness",
            Router: "/updatebusiness/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "CreateProductCart",
            Router: "/createproductcart",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "DeleteUserCart",
            Router: "/deleteusercart",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetApplyCart",
            Router: "/getapplycart",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetApprovalCart",
            Router: "/getapprovalcart",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetCart",
            Router: "/getcart",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CartController"],
        beego.ControllerComments{
            Method: "UpdateApprovalCart",
            Router: "/updateapprovalcart",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"],
        beego.ControllerComments{
            Method: "Avatar",
            Router: "/avatar/:text",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"],
        beego.ControllerComments{
            Method: "Chat",
            Router: "/chat",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ChatController"],
        beego.ControllerComments{
            Method: "HandleConnections",
            Router: "/wschat",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "ActInfo",
            Router: "/activity/actInfo",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Apply",
            Router: "/activity/apply",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/activity/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Getall",
            Router: "/activity/getall",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "HaveApply",
            Router: "/activity/haveApply",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Like",
            Router: "/activity/like",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Check",
            Router: "/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Compare",
            Router: "/check/compare",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Details",
            Router: "/check/details",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "CheckGetCheck",
            Router: "/checkgetcheck",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "CheckSignature",
            Router: "/checksignature",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Date",
            Router: "/details/date",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "MonthCheck",
            Router: "/monthcheck",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "MonthCheckSum",
            Router: "/monthchecksum",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Person",
            Router: "/person",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "SendMessage",
            Router: "/sendmessage",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:CheckController"],
        beego.ControllerComments{
            Method: "SubscribeMessage",
            Router: "/subscribemessage",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "AddWxDiary",
            Router: "/addwxdiary",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "DeleteWxDiary",
            Router: "/deletewxdiary",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "GetWxdiaries",
            Router: "/getwxdiaries",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "HtmlToDoc",
            Router: "/getwxdiaries",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "GetWxdiaries2",
            Router: "/getwxdiaries2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "GetWxDiary",
            Router: "/getwxdiary/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:DiaryController"],
        beego.ControllerComments{
            Method: "UpdateWxDiary",
            Router: "/updatewxdiary",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/get",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/search",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"],
        beego.ControllerComments{
            Method: "Tika",
            Router: "/tika",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: "/upload/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ElasticController"],
        beego.ControllerComments{
            Method: "UploadElastic",
            Router: "/uploadelastic/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "AddExcelArticles",
            Router: "/addexcelarticle/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "DeleteExcel",
            Router: "/deleteexcel/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "EditorExcelArticle",
            Router: "/editorexcelarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelArticle",
            Router: "/excelarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "ExcelCal",
            Router: "/excelcal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxExcelHistory",
            Router: "/getWxexcelhistory/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcel",
            Router: "/getexcel/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "ExcelList",
            Router: "/getexcel/excellist/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcels",
            Router: "/getexcel/getexcels/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxExcels",
            Router: "/getexcel/getwxexcels/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelCalInput",
            Router: "/getexcelcalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetexcelCalOutput",
            Router: "/getexcelcaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelHistory",
            Router: "/getexcelhistory/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelHistoryCal",
            Router: "/getexcelhistorycal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelHistoryInput",
            Router: "/getexcelhistoryinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelHistoryOutput",
            Router: "/getexcelhistoryoutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetExcelTemplePdf",
            Router: "/getexceltemplepdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetHistoryExcel",
            Router: "/gethistoryexcel/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxexcelCalInput",
            Router: "/getwxexcelcalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxExcelCalOutput",
            Router: "/getwxexcelcaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxExcelClass",
            Router: "/getwxexcelclass/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "PostExcel",
            Router: "/postexcel",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "PostWxexcel",
            Router: "/postwxexcel",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "PutexcelCalInput",
            Router: "/putexcelcalinput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "PutexcelCalOutput",
            Router: "/putexcelcaloutput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "PutExcelTemple",
            Router: "/putexceltemple",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "UpdateExcelArticle",
            Router: "/updateexcelarticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "UploadExcel",
            Router: "/uploadexcel/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "UploadExcelTemple",
            Router: "/uploadexceltemple/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ExcelController"],
        beego.ControllerComments{
            Method: "GetWxExcelArticle",
            Router: "/wxexcelarticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FileinputController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FileinputController"],
        beego.ControllerComments{
            Method: "BootstrapFileInput",
            Router: "/bootstrapfileinput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "AddWxFinance",
            Router: "/addwxfinance/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "DeleteWxFinance",
            Router: "/deletewxfinance",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "GetWxFinance",
            Router: "/getwxfinance/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "GetWxfinance2",
            Router: "/getwxfinance2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "GetWxFinanceList",
            Router: "/getwxfinancelist/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FinanceController"],
        beego.ControllerComments{
            Method: "UpdateWxFinance",
            Router: "/updatewxfinance",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowAccessContext",
            Router: "/flowaccesscontext",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowAccessContextList",
            Router: "/flowaccesscontextlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowAccessContextUpdate",
            Router: "/flowaccesscontextupdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowAction",
            Router: "/flowaction",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowActionDelete",
            Router: "/flowactiondelete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowActionList",
            Router: "/flowactionlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowActionUpdate",
            Router: "/flowactionupdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowDoc",
            Router: "/flowdoc",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowEvent",
            Router: "/flowdocevent",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowEventList",
            Router: "/flowdoceventlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowDocList",
            Router: "/flowdoclist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowDocumentDetail",
            Router: "/flowdocumentdetail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowDocumentList",
            Router: "/flowdocumentlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowDocumentList2",
            Router: "/flowdocumentlist2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroup",
            Router: "/flowgroup",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroupList",
            Router: "/flowgrouplist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroupMailbox",
            Router: "/flowgroupmailbox",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroupRole",
            Router: "/flowgrouprole",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroupRoleList",
            Router: "/flowgrouprolelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowGroupUsersList",
            Router: "/flowgroupuserslist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowNext",
            Router: "/flownext",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowNode",
            Router: "/flownode",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowNodeList",
            Router: "/flownodelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowPermission",
            Router: "/flowpermission",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowRole",
            Router: "/flowrole",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowRoleList",
            Router: "/flowrolelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowRolePermissionList",
            Router: "/flowrolepermissionlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowState",
            Router: "/flowstate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowStateDelete",
            Router: "/flowstatedelete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowStateList",
            Router: "/flowstatelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowStateUpdate",
            Router: "/flowstateupdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTransition",
            Router: "/flowtransition",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTransitionDelete",
            Router: "/flowtransitiondelete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTransitionList",
            Router: "/flowtransitionlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTransitionUpdate",
            Router: "/flowtransitionupdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowType",
            Router: "/flowtype",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTypeDelete",
            Router: "/flowtypedelete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTypeList",
            Router: "/flowtypelist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowTypeUpdate",
            Router: "/flowtypeupdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowUser",
            Router: "/flowuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowUserGroup",
            Router: "/flowusergroup",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowUserList",
            Router: "/flowuserlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowUserMailbox",
            Router: "/flowusermailbox",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowUserMailbox2",
            Router: "/flowusermailbox2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowWorkflow",
            Router: "/flowworkflow",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "FlowWorkflowList",
            Router: "/flowworkflowlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "LiuCheng",
            Router: "/liucheng",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "WorkFlow",
            Router: "/workflow",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "WxFlowDoc",
            Router: "/wxflowdoc",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "WxFlowNext",
            Router: "/wxflownext",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
        beego.ControllerComments{
            Method: "WxFlowUserMailbox2",
            Router: "/wxflowusermailbox2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlvController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlvController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlvController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlvController"],
        beego.ControllerComments{
            Method: "GetFlvList",
            Router: "/flvlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadAppreciationPhoto",
            Router: "/uploadappreciationphoto",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxAvatar",
            Router: "/uploadwxavatar",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxEditorImg",
            Router: "/uploadwxeditorimg",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxImg",
            Router: "/uploadwximg",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxImgs",
            Router: "/uploadwximgs/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxVideo",
            Router: "/uploadwxvideo/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
        beego.ControllerComments{
            Method: "UploadWxVideoCover",
            Router: "/uploadwxvideocover/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"],
        beego.ControllerComments{
            Method: "AddLocationNavigate",
            Router: "/addlocationnavigate/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"],
        beego.ControllerComments{
            Method: "AddLocationPart",
            Router: "/addlocationpart/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"],
        beego.ControllerComments{
            Method: "GetLocation",
            Router: "/getlocation/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LocationController"],
        beego.ControllerComments{
            Method: "GetLocationById",
            Router: "/getlocationbyid/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Islogin",
            Router: "/islogin",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LoginPost",
            Router: "/loginpost",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "SsoLogin",
            Router: "/ssologin",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "SsoLoginPost",
            Router: "/ssologinpost",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "WxHasSession",
            Router: "/wxhassession",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
        beego.ControllerComments{
            Method: "WxLogin",
            Router: "/wxlogin/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
        beego.ControllerComments{
            Method: "PhotoSwipe",
            Router: "/photoswipe",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
        beego.ControllerComments{
            Method: "Postdata",
            Router: "/sendmessage",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
        beego.ControllerComments{
            Method: "UserManage",
            Router: "/usermanage",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "AddMathArticles",
            Router: "/addmatharticle/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "DeleteTemple",
            Router: "/deletetemple/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "EditorMathArticle",
            Router: "/editormatharticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxHistoryTemples",
            Router: "/getWxhistorytemples/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetHistoryInput",
            Router: "/gethistoryinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetHistoryMath",
            Router: "/gethistorymath/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetHistoryMathCal",
            Router: "/gethistorymathcal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetHistoryOutput",
            Router: "/gethistoryoutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetHistoryTemples",
            Router: "/gethistorytemples/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetMath",
            Router: "/getmath/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetTemples",
            Router: "/getmath/gettemples/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxTemples",
            Router: "/getmath/getwxtemples/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "TempleList",
            Router: "/getmath/templelist/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetMathCalInput",
            Router: "/getmathcalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetMathCalOutput",
            Router: "/getmathcaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetMathTemplePdf",
            Router: "/getmathtemplepdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxMathCalInput",
            Router: "/getwxmathcalinput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxMathCalOutput",
            Router: "/getwxmathcaloutput/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxMathClass",
            Router: "/getwxmathclass/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetMathArticle",
            Router: "/matharticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "MathCal",
            Router: "/mathcal/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PostMath",
            Router: "/postmath",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PostMath2",
            Router: "/postmath2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PostMathback",
            Router: "/postmathback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PostWxMath",
            Router: "/postwxmath",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PostWxMath2",
            Router: "/postwxmath2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PutMathCalInput",
            Router: "/putmathcalinput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PutMathCalOutput",
            Router: "/putmathcaloutput",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "PutUserTemple",
            Router: "/putusertemple",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "UpdateArticle",
            Router: "/updatearticle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "UploadMathTemple",
            Router: "/uploadmathtemple/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "UploadTemple",
            Router: "/uploadtemple/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MathcadController"],
        beego.ControllerComments{
            Method: "GetWxMathArticle",
            Router: "/wxmatharticle/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"],
        beego.ControllerComments{
            Method: "CommandBuilder",
            Router: "/commandbuilder",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"],
        beego.ControllerComments{
            Method: "CommandDrop",
            Router: "/commanddrop/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"],
        beego.ControllerComments{
            Method: "CommandInfo",
            Router: "/commandinfo/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"],
        beego.ControllerComments{
            Method: "Conversion",
            Router: "/conversion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:OnlyController"],
        beego.ControllerComments{
            Method: "DownloadOnlyDoc",
            Router: "/downloadonlydoc",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "AddApplyRecharge",
            Router: "/addapplyrecharge/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "AddUserPays",
            Router: "/adduserpays",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "AddUserRecharge",
            Router: "/adduserrecharge/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "AddWxUserPays",
            Router: "/addwxuserpays",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "ApplyRecharge",
            Router: "/applyrecharge",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetApplyRecharge",
            Router: "/getapplyrecharge",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetApplyRechargeData",
            Router: "/getapplyrechargedata",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetPay",
            Router: "/getpay",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetUserGetAppreciations",
            Router: "/getusergetappreciations",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetUserMoney",
            Router: "/getusermoney",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetUserPay",
            Router: "/getuserpay",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetUserPayAppreciations",
            Router: "/getuserpayappreciations",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetUserPaylist",
            Router: "/getuserpaylist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetWxPay",
            Router: "/getwxpay",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetWxUserGetAppreciations",
            Router: "/getwxusergetappreciations",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetWxUserMoney",
            Router: "/getwxusermoney",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetWxUserPayAppreciations",
            Router: "/getwxuserpayappreciations",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PayController"],
        beego.ControllerComments{
            Method: "GetWxUserPays",
            Router: "/getwxuserpays",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"],
        beego.ControllerComments{
            Method: "AddWatermarks",
            Router: "/addwatermarks/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"],
        beego.ControllerComments{
            Method: "OnlyPdf",
            Router: "/onlypdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:PdfCpuController"],
        beego.ControllerComments{
            Method: "Test",
            Router: "/test/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProdController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProdController"],
        beego.ControllerComments{
            Method: "GetProducts",
            Router: "/project/products/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "GetProjects",
            Router: "/getprojects",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "GetProjectTree",
            Router: "/getprojecttree/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "GetWxProjects",
            Router: "/getwxprojects",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "ProjectUserRole",
            Router: "/projectuserrole",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "QuickAddWxProjTemplet",
            Router: "/quickaddwxproject",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "UserpProjectEditorTree",
            Router: "/userprojecteditortree",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ProjController"],
        beego.ControllerComments{
            Method: "UserProjectPermission",
            Router: "/userprojectpermission",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"],
        beego.ControllerComments{
            Method: "WxRegion",
            Router: "/wxregion",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"],
        beego.ControllerComments{
            Method: "WxRegist",
            Router: "/wxregist",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
        beego.ControllerComments{
            Method: "AddWxLike",
            Router: "/addwxlike/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
        beego.ControllerComments{
            Method: "AddWxRelease",
            Router: "/addwxrelease/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
        beego.ControllerComments{
            Method: "DeleteWxRelease",
            Router: "/deletewxrelease/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "GetWxPdfList",
            Router: "/getwxpdflist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchProduct",
            Router: "/searchproduct",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchProductData",
            Router: "/searchproductdata",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchProjectProduct",
            Router: "/searchprojectproduct",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchProjectProductData",
            Router: "/searchprojectproductdata",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchWxDrawings",
            Router: "/searchwxdrawings",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
        beego.ControllerComments{
            Method: "SearchWxProducts",
            Router: "/searchwxproducts",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"],
        beego.ControllerComments{
            Method: "Browse",
            Router: "/browse",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: "/detail/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"],
        beego.ControllerComments{
            Method: "Download",
            Router: "/download",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ShareController"],
        beego.ControllerComments{
            Method: "DownloadZip",
            Router: "/downloadzip",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "DownloadStandard",
            Router: "/downloadstandard/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "ElasticSearch",
            Router: "/elasticsearch",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "GetElasticStandard",
            Router: "/getelasticstandard",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/search",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "SearchWxStandards",
            Router: "/searchwxstandards",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: "/standard/upload",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "StandardPdf",
            Router: "/standardpdf",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "UploadStandard",
            Router: "/uploadstandard",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
        beego.ControllerComments{
            Method: "WxStandardPdf",
            Router: "/wxstandardpdf/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"],
        beego.ControllerComments{
            Method: "DeleteTodo",
            Router: "/deletetodo",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"],
        beego.ControllerComments{
            Method: "GetTodo",
            Router: "/gettodo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:TodoController"],
        beego.ControllerComments{
            Method: "UpdateTodo",
            Router: "/updatetodo",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: "/adduser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddWxUser",
            Router: "/addwxuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/deleteuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserByUsername",
            Router: "/getuserbyusername",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "ImportUsers",
            Router: "/importusers",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/updateuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateWxUser",
            Router: "/updatewxuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "User",
            Router: "/user/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:UserController"],
        beego.ControllerComments{
            Method: "Usermyself",
            Router: "/usermyself",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetUserVideo",
            Router: "/getuservideo/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
