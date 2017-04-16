function saveNote(){
    var hash = window.location.hash.split("/")
    if (hash.length<2) return;
    id=hash[1];

    $.post( "api/modify-note/"+id, { 'note': {
        'title': $("#noteform-title").val().replace(/\n/g, "\\n"),
        'notebook': $("#noteform-notebook").val(),
        'body': $("#noteform-body").val().replace(/\n/g, "\\n")
    } } );

}


function deleteNotebook(id){
    var resp = confirm ("Are you sure you want to remove this notebook with all child notes?")
    if (!resp) return;
    resp = prompt ("Again: Are you sure you want to remove this notebook with all child notes?\n\nPlease enter 'sudo delete'")
    if (resp!="sudo delete") return;
    $.get("api/delete-notebook/"+id);
    alert("deleted!")
}
function renameNotebook(id){
    var resp = prompt ("Please give this notebook ("+$.grep(database.notebooks,function(n,i){return n.id==id})[0].title+") new name:")
    if (resp==false || resp=="") return

    return;
    $.post( "api/modify-notebook/"+id, { 'note': {
        'title': $("#noteform-title").val().replace(/\n/g, "\\n"),
        'notebook': $("#noteform-notebook").val(),
        'body': $("#noteform-body").val().replace(/\n/g, "\\n")
    } } );
}