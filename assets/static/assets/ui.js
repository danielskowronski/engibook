function updateTitle(title, descr){
    $("#title").html(title)
    $("#descr").html(descr)
}

function filterByNotebook(id){
    window.location.hash = "#notes"
    $("#content").html();
    if (id>=0){
        updateTitle(notebooksMap[id],"")
        $(".noteEntry").hide()
        $(".noteEntry[data-notebook='"+id+"']").show()
    }
    else {
        updateTitle("%all% notes index","")
        $(".noteEntry").show()
    }
}
function filterByString(serach){
    $(".noteEntry").hide();
    $.each( $(".noteEntry"), function( key, val ) {
        var body = $(val).html()
        if ($("#regexSearchEnabled").prop('checked')){
            if (body.match(serach)){
                $(val).show()
            }
        }       
        else {
            if (body.indexOf(serach)!=-1){
                $(val).show()
            }
        }
    });
}
function expandOrCollapseNote(caller){
    var target = $(caller).parent();
    if (!target.prop("data-expanded") || target.prop("data-expanded") =="false"){
        target.prop("data-expanded","true")
        target.css("max-height","none")
    }
    else {
        target.prop("data-expanded","false")
        target.css("max-height","50px")
    }
}