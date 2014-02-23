// use ring code
$(function(){
    var pageLoadTime = new Date().getTime();
    function reloadIfSourceChanged() {
        $.get('/@reveltang_dev/__source_changed?clientSince=' + pageLoadTime, function(ret){
            if(ret == 'true'){
                window.location.reload();
            }else{
                setTimeout(reloadIfSourceChanged, 1000);
            }
        }).fail(function() {
            window.location.reload();
        });
    }
    reloadIfSourceChanged();
});
