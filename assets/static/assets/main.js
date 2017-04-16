var database = {notebooks:[],notes:[]}; //locally laoded database
var notebooksMap = {};

var converter = new showdown.Converter();
converter.setOption('simplifiedAutoLink', 'true');
converter.setOption('strikethrough', 'true');
converter.setOption('simpleLineBreaks', 'true');
showdown.setFlavor('github');

$(function() { 
    pageReloaded(); 
    window.onhashchange = pageReloaded;

    $("form").bind("input change propertychange", function (evt) {
        if (window.event && event.type == "propertychange" && event.propertyName != "value")
            return;
        window.clearTimeout($(this).data("timeout"));
        $(this).data("timeout", setTimeout(function () {
            filterByString($('#searchForm').val())
        }, 100));
    });

    $("#noteform-body").bind("input change propertychange", function (evt) {
        if (window.event && event.type == "propertychange" && event.propertyName != "value")
            return;
        window.clearTimeout($(this).data("timeout"));
        $(this).data("timeout", setTimeout(function () {
            $("#noteform-preview").html(converter.makeHtml($("#noteform-body").val()))
        }, 100));
    });
});