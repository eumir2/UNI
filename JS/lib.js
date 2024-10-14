const apiKey = "api_key=2461e2dc84a8e379ad6a8309c585a25d";
const lang = "language=it-IT";
const baseUrl = "https://api.themoviedb.org/3";

function getPage() {
    query = window.location.search;
    var page = new URLSearchParams(query).get('page');

    if (page == null || page < 1) {
        document.getElementById('page-prev').innerHTML = 1
        document.getElementById('page-current').innerHTML = 2
        document.getElementById('page-next').innerHTML = 3

        return 1;
    }
    page = parseInt(page);
    document.getElementById('page-prev').innerHTML = page - 1
    document.getElementById('page-current').innerHTML = page
    document.getElementById('page-next').innerHTML = page + 1
    return page;
}

function setPage(page) {
    query = window.location.search;
    var search = new URLSearchParams(query);
    search.set('page', page);
    window.location.search = search.toString();
}

function getCast(id){
    url = baseUrl + "/movie/" + id + "/credits?" + apiKey + "&" + lang;
    fetch(url)
        .then(response => response.json())
        .then(cast => schedaCast(cast));
}

function getSchedaFilm(id,type="movie") {
    url = baseUrl + "/"+ type + "/" + id + "?" + apiKey + "&" + lang;
    fetch(url)
        .then(response => response.json())
        .then(film => scheda(film,type));
}



function getSchedaAttore(id) {
    url = baseUrl + "/person/" + id + "?" + apiKey + "&" + lang;
    fetch(url)
        .then(response => response.json())
        .then(attore => schedaAttore(attore));
}


function schedaCast(cast){
    document.getElementById('cast').innerHTML  = ""
    for(var i=0; i < cast.cast.length; i++){
        document.getElementById('cast').innerHTML 
        += "<a href='scheda-attore.html?id=" + cast.cast[i].id + "'>" + cast.cast[i].name + "</a><br>"

    }

 

}


function schedaAttore(attore) {

    document.getElementsByClassName('card-title')[0].innerHTML
        = attore.name;

    document.getElementsByClassName('card-text')[0].innerHTML
        = attore.biography;

    document.getElementsByClassName('card-img-top')[0].src =
        "https://image.tmdb.org/t/p/w200" + attore.profile_path;
    document.getElementsByClassName('list-group-item')[1].innerHTML = ""
    for(var i=0; i < attore.also_known_as.length; i++){
        document.getElementsByClassName('list-group-item')[1].innerHTML 
        += attore.also_known_as[i] + "<br>"

    }
    document.getElementsByClassName('list-group-item')[0].innerHTML = "<a href='" + attore.homepage + "'>Sito web</a>"
    document.getElementById('voti').style.width = parseFloat(film.popularity)  + "%"


}

function scheda(film,type="movie") {
    console.log(film.media_type)
    if(type == 'movie'){
        document.getElementsByClassName('card-title')[0].innerHTML
        = film.title;
    }else{
        document.getElementsByClassName('card-title')[0].innerHTML
        = film.name;
    }


    document.getElementsByClassName('card-text')[0].innerHTML
        = film.overview;

    document.getElementsByClassName('card-img-top')[0].src =
        "https://image.tmdb.org/t/p/w200" + film.poster_path;
    document.getElementsByClassName('list-group-item')[1].innerHTML = ""
    for(var i=0; i < film.genres.length; i++){
        document.getElementsByClassName('list-group-item')[1].innerHTML 
        += film.genres[i].name + "<br>"

    }

    // document.getElementsByClassName('list-group-item')[1].innerHTML =
    // film.vote_average + " (" + film.vote_count + ") <br>"
    document.getElementById('voti').style.width = parseFloat(film.vote_average) * 10 + "%"
    document.getElementById('voti').innerHTML = parseFloat(film.vote_average) * 10 + "%"

}

function getPopular(page = 1) {

    url = baseUrl + "/movie/popular?" + apiKey + "&" + lang + "&page=" + page
    fetch(url)
        .then(response => response.json())
        .then(popular => cards(popular));

}

function cerca(value){
    getSearch(value)
}
function cercaMulti(value){
    getMultiSearch(value)
}
function cercaFilm(value){
    getSearch(value)
}

function getSearch(query,page = 1) {

    url = baseUrl + "/search/movie?" + apiKey + "&" + lang + "&page=" + page + "&query=" + query
    console.log(url)
    fetch(url)
        .then(response => response.json())
        .then(searchResult => cards(searchResult));

}
function getMultiSearch(query,page = 1) {

    url = baseUrl + "/search/multi?" + apiKey + "&" + lang + "&page=" + page + "&query=" + query
    console.log(url)
    fetch(url)
        .then(response => response.json())
        .then(searchResult => cards(searchResult));

}

function cards(popolari) {
    var card = document.getElementById('card-film');
    for (var i = popolari.results.length - 1; i >= 0; i--) {
        // for (var i =0; i <  popolari.results.length;   i++) {

        var clone = card.cloneNode(true);
        clone.id = 'badge-' + i;

        if(popolari.results[i].media_type == "movie"){
            
            clone.getElementsByClassName('card-title')[0].innerHTML
                = popolari.results[i].title;
        }else{
          
            clone.getElementsByClassName('card-title')[0].innerHTML
                = popolari.results[i].name;
        }

        clone.getElementsByClassName('card-text')[0].innerHTML
            = popolari.results[i].overview.substring(0, 200) + "...";
        
        clone.getElementsByClassName('card-img-top')[0].src =
            "https://image.tmdb.org/t/p/w200" + popolari.results[i].poster_path;
        card.after(clone);

        clone.getElementsByClassName('btn-primary')[0].href = clone.getElementsByClassName('btn-primary')[0].href + popolari.results[i].id + "&type=" + popolari.results[i].media_type

        clone.classList.remove('d-none');
    }


}

function erroreImmagine(img){
    console.log(img)
    img.src = "https://via.placeholder.com/300x400?text=Immagine"

}
