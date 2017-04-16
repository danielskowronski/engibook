function pageReloaded(){
    $("#noteform").hide()
    $("#notebookform").hide()

    database = {notebooks:[],notes:[]};
    $("#content").html("")
    $("#noteform-notebook").html("")

	var action = window.location.hash.split("/")[0];
    switch (action){
        case "#trash"://TODO 
            updateTitle("trash - not implemented","shows all notes")
            break;
        case "#editNote":
            updateTitle("vi note","add/edit note")
            $("#content").html("")
            $("#noteform").show()
            if (window.location.hash.split("/")[1]!="new"){
                id=window.location.hash.split("/")[1];
                $("#noteform-dellink").show()
                $("#noteform-dellink").prop("href", "api/delete-note/"+id)
                setTimeout(function(){
                    $.getJSON( "api/get-one/"+id, function( data ) {
                        $("#noteform-title").val(data.title)
                        $("#noteform-body").val(data.body)
                        $("#noteform-notebook").val(data.notebook)
                        $("#noteform-preview").html(converter.makeHtml($("#noteform-body").val()))
                    });
                }, 250)
            } else {
                $("#noteform-dellink").hide()
            }

            break;
        case "#notebooks": //TODO: invent better route name
            updateTitle("vi notebooks","manages notebooks")
            $("#content").html("")
            $("#notebookform").show()
            break;    
        case "#man": //TODO: invent better route name
            updateTitle("man engibook","basic help")
            $("#content").html("notes (markdown/text entities with titles) belongs to notebooks<br />you can filter notes by notebooks or search by substring or regex")
            break;              
        default:
            window.location.hash = "#notes"
            filterByNotebook(-1);
            break;
    }

    Pace.start();
	$.getJSON( "api/data.json", function( data ) {
        database.notebooks = data.notebooks;
        $("#notebooks").html("")
        $("#notebookform-list").html("")
        $("#notebooks").append('<li class="pure-menu-item"><a href="javascript:filterByNotebook(-1)" class="pure-menu-link">%all%</a></li>')
        $.each( database.notebooks, function( key, val ) {
            $("#notebooks").append('<li class="pure-menu-item"><a href="javascript:filterByNotebook('+val.id+')"  class="pure-menu-link">'+val.title+'</a></li>')
            $("#noteform-notebook").append('<option value='+val.id+'>'+val.title+'</option>')
            $("#notebookform-list").append('<div class="notebookEditor">'+val.title+' <a onClick="renameNotebook('+val.id+')">[rename]</a> <a style="color: red" onClick="deleteNotebook('+val.id+')">[delete]</a> </div>')

        });
        
        database.notes = data.notes;
        if (window.location.hash == "#notes") {
            $.each( database.notes, function( key, val ) {
                var notebookName = (database.notebooks[val.notebook]).title
                $("#content").append('<div class="noteEntry" data-notebook="'+val.notebook+'">'+
                    '<div class="noteTitle" onClick="expandOrCollapseNote(this)"> ['+notebookName+"] "+val.title+
                    '<a href="#editNote/'+val.id+'">[edit note]</a></div>'+
                    '<div class="noteBody">'+converter.makeHtml(val.body)+'</div>'+
                    '</div>'
                )
            });
        }

        Pace.stop();
	});
    
}