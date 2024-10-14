
function controllaEmail(email) {
    //console.log(email)
    if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)) {
        return true;
        //console.log("mail giusta");
    } else {
        return false;
    }
};

function confermaEmail() {
    //console.log(email2.value)
    var email1 = document.getElementById("email").value;
    var email2 = document.getElementById("emailC").value;
    //console.log(mail1);
    if (email1 == email2) {
        //console.log("mail uguali")
        document.getElementById("emailAlert2").classList.add("d-none");
        return true;
    } else {
        document.getElementById("emailAlert2").classList.remove("d-none");
        return false;
    }
};

function controllaPassword(password) {
    //console.log(password)
    if (/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/.test(password)) {
        //console.log("password giusta")
        return true;
    } else {
        return false;
    }
}

function confermaPassword() {
    let password1 = document.getElementById("password").value;
    let password2 = document.getElementById("passwordC").value;
    if (password1 == password2) {
        //console.log("password coincidono");
        document.getElementById("pwAlert2").classList.toggle("d-none");
        return true;
    } else {
        document.getElementById("pwAlert2").classList.remove("d-none");
        return false;
    }
}

//funzione che controlla che se un utente è presente o meno all'interno del local storage
function controllaEsistenza(newUtente, utenti) {
    //console.log(Array.isArray(utenti));
    var risultato = utenti.find(utente => utente.email == newUtente.email);

    if (risultato != undefined) {
        return true;
    } else {
        return false
    }
}





const PUBLIC_KEY = "a2a602917a527ed53e192601aa7c4138";
const PRIVATE_KEY = "738c8b65eecd53027d1366b21bb4ca6b35bb5203";
const baseUrl = "https://gateway.marvel.com"; //URL base
const timestamp = Date.now();

function calculateMD5(num, str1, str2) {
    const input = `${num}${str1}${str2}`;
    const md5Hash = CryptoJS.MD5(input);
    return md5Hash.toString();
}

//funzione che carica tutti dati relativi all'utente
function caricaUtente() {
    var url = new URLSearchParams(window.location.search);
    nomeUtente = url.get("name");
    email = url.get("email");
    utente = getUtente(email);
    //console.log(utente);

    document.getElementById("utente").innerHTML = "CIAO," + nomeUtente.toUpperCase() + "!";
    var profilo = "profilo/";

    switch(utente.supereroe){
        case "Iron Man" : 
            profilo += "iron_man.svg";
            break;
        case "Spiderman" : 
            profilo += "spiderman.svg";
            break;
        case "Thor" : 
            profilo += "thor.svg";
            break;
        case "Hulk" : 
            profilo += "hulk.svg";
            break;
        case "Capitan America" : 
            profilo += "capitan_america.svg";
            break;
    }

    document.getElementById("avatar").src = profilo;
  }

//NAVIGAZIONE DEL MENU DROPDOWN PROFILO--------------------
  function modificaUtente() {
    window.location.href = "Profilo.html?name=" + nomeUtente + "&email=" + email + "&modifica=" + true;
  }

  function vislualizzaUtente() {
    window.location.href = "Profilo.html?name=" + nomeUtente + "&email=" + email + "&modifica=" + false;
  }

  function confermaElimina() {
    document.getElementById("prompt-eliminazione").classList.remove("d-none");
  }

  function elimina() {
    var utenti = []

    if (window.localStorage.getItem('utenti') != null) {
      utenti = JSON.parse(window.localStorage.getItem('utenti'))
    }

    var index = utenti.findIndex(utente => utente.email == email);
    utenti.splice(index, 1);
    //console.log(index);
    //console.log(utenti[index]);
    window.localStorage.setItem('utenti', JSON.stringify(utenti))

    document.getElementById("prompt-eliminazione").classList.add("d-none");
    window.location.href = "Welcome.html";
  }

  function annulla() {
    document.getElementById("prompt-eliminazione").classList.add("d-none");
  }

//------------------------------------------------

//funzione che scarica i dati relativi ai primi 99 eroi nel DB marvel
function getEroi() {
    const hash = calculateMD5(timestamp, PRIVATE_KEY, PUBLIC_KEY);
    url = baseUrl + "/v1/public/characters?limit=99&apikey=" + PUBLIC_KEY + "&ts=" + timestamp + "&hash=" + hash;
    //console.log(url);
    fetch(url)
        .then(response => response.json())
        .then(eroi => salvaEroiNelLocalStorage(eroi.data.results))
        .catch(error => window.alert(error))

    var tmp = JSON.parse(window.localStorage.getItem("cards"));
    creaCarte(tmp, 0);
}

//funzione che salva gli eroi scaricati nel localstorage
function salvaEroiNelLocalStorage(vettore) {
    localStorage.setItem("cards", JSON.stringify(vettore));
}


//funzione che crea crea le cards e le visualizza sulla pagina (il flag serve a far capire se devo mostrare la collezione dell'utente, le carte totali o gli scambi)
//flag = 0 : tutte le carte
//flag  = 1 : collezione dell'utente
//flag = 2 : scambi
function creaCarte(tmp, flag) {

    //console.table(tmp);
    var card = document.getElementById("card-hero");
    var cards = JSON.parse(window.localStorage.getItem("cards"));
    var scambi = JSON.parse(window.localStorage.getItem('scambi'));

    var url = new URLSearchParams(window.location.search);
    email = url.get("email");
    utente = getUtente(email);

    //console.log(utente);
    for (let i = 0; i < tmp.length; i++) {
        if (tmp[i] != null) {
            var clone = card.cloneNode(true);
            //console.log(tmp[i].id);

            if (flag == 2) {
                if (scambi[i].email == utente.email) {
                    continue;
                }
                clone.getElementsByClassName("card-text")[1].innerHTML = "DISPONIBILITA: " + scambi[i].copie;
                clone.getElementsByClassName("card-text")[2].innerHTML = "OFFERTO DA : " + scambi[i].email;

                clone.getElementsByClassName("card-footer")[1].value = scambi[i].email + "-" + scambi[i].pos + "-" + scambi[i].idCarta + "-" + scambi[i].copie;
            }

            clone.getElementsByClassName("card-footer")[0].value = tmp[i].id;
            clone.getElementsByClassName("card-img-top")[0].src =
                tmp[i].thumbnail.path + "/portrait_uncanny." + tmp[i].thumbnail.extension;

            clone.getElementsByClassName("card-title")[0].innerHTML = "NOME: " +
                tmp[i].name.toUpperCase();

            clone.getElementsByClassName("card-text")[0].innerHTML = "DESCRIZIONE: " +
                tmp[i].description.substring(0, 50) + "...";
            //console.log(tmp[i].description)


            if (flag == 1) {
                //console.log(t);
                var index = cards.findIndex(c => c.id == tmp[i].id);
                //console.log(index);
                if (utente.carte[index] != null) {
                    clone.getElementsByClassName("card-text")[1].innerHTML = "DISPONIBILITA: " + utente.carte[index];
                    clone.getElementsByClassName("card-footer")[1].value = email + "-" + index + "-" + tmp[i].id + "-" + utente.carte[index];
                }
            }
            card.after(clone)
            clone.classList.remove("d-none");
        }
    }
}

//funzione che ritorna un utente data la sua email
function getUtente(email) {
    if (window.localStorage.getItem('utenti') != null) {
        var utenti = JSON.parse(window.localStorage.getItem('utenti'))
    } else {
        return false;
    }

    return utenti.find(utente => utente.email == email);
}


//estrapola carica la pagina dell'eroe cliccato inviando il suo id
function caricaSchedaEroe(id) {
    //console.log(id);
    window.location.href = "SchedaEroe.html?id=" + id;
}

//funzione che estrapola l'eroe dal localstorage tramite il suo id preso dall'url della pagina
//e carica le sue informazioni
function loadEroe() {
    let url = new URLSearchParams(window.location.search);
    var id = url.get("id");
    //console.log(id)
    let eroi = JSON.parse(window.localStorage.getItem("cards"));
    //console.log(eroi)
    let eroe = eroi.find(tmp => tmp.id == id);
    //console.log(eroe)

    document.getElementById("immagineEroe").src =
        eroe.thumbnail.path + "/portrait_uncanny." + eroe.thumbnail.extension;

    document.getElementById("nomeEroe").innerHTML = "Nome: " + eroe.name

    document.getElementById("descrizioneEroe").innerHTML = "Descrizione: " + eroe.description;

    getSeries(eroe.series.collectionURI);
    getEvents(eroe.events.collectionURI);
    getComics(eroe.comics.collectionURI);
}


//funzione che scarica le serie dove compare un determinato supereroe
function getSeries(uri) {
    const hash = calculateMD5(timestamp, PRIVATE_KEY, PUBLIC_KEY);
    var url = uri + "?&apikey=" + PUBLIC_KEY + "&ts=" + timestamp + "&hash=" + hash;
    //console.log(url);
    fetch(url)
        .then(response => response.json())
        .then(tmp => loadData(tmp.data.results, "#caroselloSerie"))
        .catch(error => console.log('error', error))

    //console.log(series);
}

//funzione che scarica gli eventi dove compare un determinato supereroe
function getEvents(uri) {
    const hash = calculateMD5(timestamp, PRIVATE_KEY, PUBLIC_KEY);
    var url = uri + "?&apikey=" + PUBLIC_KEY + "&ts=" + timestamp + "&hash=" + hash;
    //console.log(url);
    fetch(url)
        .then(response => response.json())
        .then(tmp => loadData(tmp.data.results, "#caroselloEventi"))
        .catch(error => console.log('error', error))

    //console.log(series);
}

//funzione che scarica i comics dove compare un determinato supereroe
function getComics(uri) {
    const hash = calculateMD5(timestamp, PRIVATE_KEY, PUBLIC_KEY);
    var url = uri + "?&apikey=" + PUBLIC_KEY + "&ts=" + timestamp + "&hash=" + hash;
    //console.log(url);
    fetch(url)
        .then(response => response.json())
        .then(tmp => loadData(tmp.data.results, "#caroselloComics"))
        .catch(error => console.log('error', error))

    //console.log(series);
}



function loadData(tmp, id) {
    var carouselInner = document.querySelector(id);
    if (tmp.length == 0) {
        carouselInner.innerHTML = "Nessun elemento disponibile";
        return;
    }
    for (let i = 0; i < tmp.length; i++) {
        var carouselItem = document.createElement('div');

        carouselItem.classList.add('carousel-item');
        carouselItem.classList.add('active');

        var img = document.createElement('img');
        img.src = tmp[i].thumbnail.path + "/portrait_fantastic." + tmp[i].thumbnail.extension;
        var nome = document.createElement('div');
        nome.innerHTML = tmp[i].title.substring(0, 60) + "..."
        nome.classList.add('text');

        carouselItem.appendChild(img);
        carouselItem.appendChild(nome);
        carouselInner.appendChild(carouselItem);

    }

    //console.log(carouselInner);
}

//funzione che verifica che l'utente abbia abbastanza credito per poter acquistare un pacchetto
function controllaAcquisto(utente) {
    document.getElementById('acquisto').setAttribute('disabled', true);
    if (utente.credito > 0) {
        var users = JSON.parse(window.localStorage.getItem('utenti'))
        var newUtente = acquistaPachetto(utente);
        for (let i = 0; i < users.length; i++) {
            if (users[i].email == newUtente.email) {
                users[i] = newUtente;
                break;
            }
        }
        //console.table(users);

        window.localStorage.setItem("utenti", JSON.stringify(users));
        //acquistaPachetto()
    } else {
        window.alert("Credito Insufficiente");
        return;
    }
}

//funzione che aggiorna il credito e le carte ottenute dell'utente dopo
//l'apertura del pacchetto
function acquistaPachetto(utente) {
    utente.credito--;
    //console.table(utente);

    var cards = JSON.parse(window.localStorage.getItem('cards'));
    var cardsTmp = [];
    for (let i = 0; i < 5; i++) {
        let tmp = getRandomInt(0, 99);


        //console.log(id);
        cardsTmp.push(cards[tmp]);

        if (isNaN(utente.carte[tmp])) {
            utente.carte[tmp] = 1;
        } else {
            utente.carte[tmp]++;
        }
    }

    //console.table(cardsTmp);
    creaCarte(cardsTmp, 0);
    document.getElementById('container-pacchetto').classList.remove("d-none");
    return utente;

}

//funzioni di navigazione menu-----------
function apriCarte() {
    window.location.href = "Home.html?name=" + nomeUtente + "&email=" + email;
}

function apriNegozio() {
    window.location.href = "Negozio.html?name=" + nomeUtente + "&email=" + email;
}

function apriCollezione() {
    window.location.href = "Collezione.html?name=" + nomeUtente + "&email=" + email;
}

function apriScambi() {
    window.location.href = "Scambi.html?name=" + nomeUtente + "&email=" + email;
}
//------------------------------------

//funzione che fa 'comparire' il form di acquisto crediti
function acquistaCrediti() {
    document.getElementById("form-pagamento").classList.remove("d-none");
}

//funzione che aggiona il prezzo ogni volta che viene cambiato il numero di crediti che si vogliono acquistare
function aggiornaPrezzo(prezzo) {
    //window.alert(prezzo);
    let p = document.getElementById('crediti').value;
    document.getElementById('prezzo').value = p + "$";
}

//funzione che 'compra i crediti per l'utente'
function compraCrediti(email) {
    var tmp = document.getElementById('prezzo').value;
    var crediti = parseInt(tmp);
    //window.alert(mail);
    utente = getUtente(email);
    utente.credito += crediti;
    document.getElementById("form-pagamento").classList.add("d-none");
    //console.log(utente);

    //reinserisco l'utente nel local storage
    var users = JSON.parse(window.localStorage.getItem('utenti'))
    for (let i = 0; i < users.length; i++) {
        if (users[i].email == utente.email) {
            users[i] = utente;
            break;
        }
    }
    //console.table(users);

    window.localStorage.setItem("utenti", JSON.stringify(users));
    apriNegozio();    
}

//funzione che mostra la collezione dell'utente
function loadCollezione(utente) {
    //console.log(utente);   
    var tmp = JSON.parse(window.localStorage.getItem('cards'));
    var cards = [];
    for (let i = 0; i < utente.carte.length; i++) {
        if (utente.carte[i] != null) {
            cards.push(tmp[i]);
        }
    }
    creaCarte(cards, 1);
    //console.log(cards);    
}

//funzione che mostra all'utente quante 'copie' poter scambiare
function selezionaCopie(tmp) {
    //window.alert(tmp);
    var params = tmp.split("-");
    //console.log(params);
    var d = parseInt(params[3]);
    if (d <= 1) {
        alert("Non puoi scambiare la carta con una sola copia!");
        return;
    }
    document.getElementById("disponibilita").innerHTML = "DISPONIBILITA': " + d;
    document.getElementById("copie").setAttribute("max", d - 1);
    document.getElementById("copie").setAttribute("min", 1);
    document.getElementById("copie").setAttribute("value", 1);
    //let t = document.getElementById("copie").value;
    //console.log(t);
    document.getElementById("params").value = tmp;
    document.getElementById("prompt-copie").classList.remove("d-none");
}

//funzione che crea lo scambio e lo mette nel local storage
function creaScambio() {
    var copie = document.getElementById("copie").value;
    //console.log(copie);
    var params = document.getElementById("params").value.split("-")
    var utente = getUtente(params[0]);


    //aggiorno le carte disponibili dell'utente
    var utenti = JSON.parse(window.localStorage.getItem('utenti'));
    let i = utenti.findIndex(d => d.email == utente.email);
    let tmp = utenti[i];
    tmp.carte[params[1]] -= parseInt(copie);
    tmp.scambiDisponibili += parseInt(copie);
    utenti[i] = tmp;
    window.localStorage.setItem('utenti', JSON.stringify(utenti));

    var scambi = [];
    //console.log(params)
    var scambio = {
        'email': params[0],
        'pos': params[1],
        'idCarta': params[2],
        'copie': copie
    }
    //console.log(window.localStorage.getItem('scambi'));

    if (window.localStorage.getItem('scambi') !== null) {
        scambi = JSON.parse(window.localStorage.getItem('scambi'));
    }

    if (!checkScambio(scambio, scambi)) {
        console.log("falso");
        scambi.push(scambio);
        window.localStorage.setItem('scambi', JSON.stringify(scambi));
    }

    apriCollezione();
}

//funzione che controlla se lo scambio non sia gia presente all'interno del local storage, se è presente verrà aggiornato lo scambio
function checkScambio(scambio, scambi) {
    //console.log(scambi);
    var risultato = scambi.filter((t) => t.email === scambio.email && t.idCarta === scambio.idCarta);
    //console.log(risultato);
    if (risultato.length > 0) {
        let i = scambi.findIndex(d => d.email == scambio.email && d.idCarta == scambio.idCarta);
        let tmp = scambi[i];
        let v1 = parseInt(tmp.copie);
        let v2 = parseInt(scambio.copie);
        tmp.copie = v1 + v2;
        scambi[i] = tmp;
        window.localStorage.setItem('scambi', JSON.stringify(scambi));
        return true;
    } else {
        return false;
    }
}

//funzione che mostra gli scambi disponibili
function loadScambi() {
    var scambi = JSON.parse(window.localStorage.getItem('scambi'));
    var cards = JSON.parse(window.localStorage.getItem('cards'));

    var mercato = [];

    for (let i = 0; i < scambi.length; i++) {
        mercato.push(cards[scambi[i].pos]);
    }
    //console.log(mercato)
    //console.log(email);

    let utente = getUtente(email);
    document.getElementById('scambi').innerHTML = "SCAMBI DISPONIBILI: " + utente.scambiDisponibili;
    creaCarte(mercato, 2);
}


//funzione che permette all'utente di selezionare il numero di copie da ottenere in base al numero di vendite create o al numero di copie disponibili
function selezionaScambio(tmp) {
    let params = tmp.split("-");
    let utente = getUtente(email);
    let sC = utente.scambiDisponibili;


    //se l'utente non ha ancora 'messo in vendita niente'
    if (sC == 0) {
        window.alert("Non puoi barattare se non proponi prima degli scambi");
        return;
    }

    //trova il massimo tra sC e params[3] 
    let max = Math.min(sC, params[3]);
    //console.log(max);

    document.getElementById("copie").setAttribute("max", max);
    document.getElementById("copie").setAttribute("min", 1);
    document.getElementById("copie").setAttribute("value", 1);

    document.getElementById("params").value = tmp;
    document.getElementById("prompt-scambi").classList.remove("d-none");
}

//funzione che effettua lo scambio e aggiorna i dati nel localstorage
function effettuaScambio(){
    var scambi = JSON.parse(window.localStorage.getItem('scambi')); 
    var utenti = JSON.parse(window.localStorage.getItem('utenti'));
    var copie = parseInt(document.getElementById("copie").value);
    let tmp = document.getElementById("params").value
    let params = tmp.split("-");

    //console.log(copie);

    //aggiorno le copie disponibil nello scambio
    let i = scambi.findIndex(d => d.email == params[0] && d.idCarta == params[2]);
    let s = scambi[i];
    let v1 = parseInt(s.copie)
    let v = v1 - copie;
    //console.log(v);
    s.copie = v;
    scambi[i] = s;
    
    //tolgo gli scambi con copie = 0
    let newScambi = scambi.filter((t) => t.copie > 0);
    


    //aggiorno la quantità di scambi disponibili dell'utente
    let j = utenti.findIndex(d => d.email == email);
    let u = utenti[j];
    //console.log(u);
    v = parseInt(u.scambiDisponibili) - copie;
    u.scambiDisponibili = v;
    //aggiorno la lista delle carte dell'utente
    u.carte[params[1]] += copie;
    utenti[j] = u;

    //inserisco i dati nel localstorage
    window.localStorage.setItem('scambi',JSON.stringify(newScambi));
    window.localStorage.setItem('utenti',JSON.stringify(utenti));

    apriScambi();
}

function getRandomInt(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}



