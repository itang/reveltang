// use ring code
$(function(){
    var pageLoadTime = new Date().getTime();
    function reloadIfSourceChanged() {
        $.get('/@dev/__source_changed?clientSince=' + pageLoadTime, function(ret){
            if(ret == 'true'){
                window.location.reload();
            }else{
                setTimeout(reloadIfSourceChanged, 500);
            }
        }).fail(function() {
            window.location.reload();
        });
    }
    reloadIfSourceChanged();
});
