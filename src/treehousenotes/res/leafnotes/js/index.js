function onLoadScripts(){
    DocumentOnLoadNavBar();
    initControls();
    BindMouseEvents();
    
    //last one to calculate.
    calculateFrameDivSize(true);
}