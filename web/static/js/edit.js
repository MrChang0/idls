let editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    mode: "lua",
    lineNumbers: true,    //显示行号
    theme: "dracula",    //设置主题
    lineWrapping: true,    //代码折叠
    foldGutter: true,
    gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"],
    indentUnit: 4,
});

$("#commit").click(function () {
    const uuid = $("#uuid").val();
    const code = editor.getValue();
    const url = "/manger/" + uuid + "/code";
    $.post(url,{uuid:uuid,code:code},function (result) {
        if(result.statue === "success"){
            $("#err").val("无错误")
        }else{
            $("#err").val(result.err)
        }
    })
});