let collapsed = true;

var collapsedCoauthors = document.getElementById("collapsedCoauthors");
if (collapsedCoauthors){
    collapsedCoauthors.addEventListener("click", displayCoauthors);
}

var uncollapsedAction = document.getElementById("uncollapsedAction");
if (uncollapsedAction){
    uncollapsedAction.addEventListener("click", hideCoauthors);
}

function displayCoauthors(){
    document.getElementById("uncollapsedCoauthors").classList.remove("hidden");
}

function hideCoauthors(){
    document.getElementById("uncollapsedCoauthors").classList.add("hidden");
}