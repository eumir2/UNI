<?php
session_start();

//funzione per applicare le modifiche
function update($query, $conn){
    //var_dump($query);

   $result = pg_query($conn, $query);
    if (!$result) {
        echo "Si è verificato un errore.<br/>";
        echo pg_last_error($conn);
        echo ("<br><br>Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
    exit();
    }
    else {
    //$cmdtuples = pg_affected_rows($result);
        echo "Aggiornamento avvenuto con successo<br><a href='inserimento_e_modifica.php'>ritorna</a>";
    };
}

//var_dump($_POST);
if (isset($_POST['update']) &&isset($_SESSION['tbl'])) {//sono stati passati correttamente i dati
    $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
    if (!$conn) {
      echo 'Connessione al database fallita.';
      exit();
    }else{
        $tbl = $_SESSION['tbl'];
        //var_dump($_POST);
        switch($tbl){
            case 'piscina' :
                $nomepiscina = isset($_POST['nome']) ? $_POST['nome'] : NULL;
                $via = isset($_POST['via']) ? $_POST['via'] : NULL;
                $nciv = (int)(isset($_POST['nciv']) and is_numeric($_POST['nciv'])) ? $_POST['nciv'] : NULL;
                $nvasche = (int)(isset($_POST['nvasche']) and is_numeric($_POST['nvasche'])) ? $_POST['nvasche'] : NULL;
                $inizio = isset($_POST['inizio']) ? $_POST['inizio'] : NULL;
                $fine = isset($_POST['fine']) ? $_POST['fine'] : NULL;

                //attributi responsabile
                $cfres = isset($_POST['cfres']) ? $_POST['cfres'] : NULL;


                $query = "UPDATE piscina SET 
                    nome = '$nomepiscina',
                    cfresponsabile = '$cfres',
                    via = '$via',
                    n_civ = '$nciv',
                    apertura = '$inizio',
                    chiusura = '$fine',
                    nvasche = '$nvasche'
                    WHERE nome = '$nomepiscina';";
                    //var_dump($query);
                break;
            case 'istruttore' :
                $numerobadge = isset($_POST['nbadge']) ? $_POST['nbadge'] : NULL;
                $nome = isset($_POST['nomei']) ? $_POST['nomei'] : NULL;
                $cognome = isset($_POST['cognomei']) ? $_POST['cognomei'] : NULL;
                $nascita = (isset($_POST['datai'])) ? $_POST['datai'] : NULL;
                //$tel = isset($_POST['teli']) ? $_POST['teli'] : NULL;

                $query = "UPDATE istruttore SET
                    numerobadge = '$numerobadge',
                    nome = '$nome',
                    cognome = '$cognome',
                    datadinascita = '$nascita'
                    WHERE numerobadge = '$numerobadge'";

                //var_dump($query);
                //update($query, $conn);
                break;
            case 'corso' : 
                $codicecorso = isset($_POST['codicecorso']) ? $_POST['codicecorso'] : NULL;
                $nomecorso = isset($_POST['ncorso']) ? $_POST['ncorso'] : NULL;
                $minpar= isset($_POST['minpar']) ? $_POST['minpar'] : NULL;
                $maxpar = isset($_POST['maxpar']) ? $_POST['maxpar'] : NULL;
                $nlezioni = isset($_POST['nlezioni']) ? $_POST['nlezioni'] : NULL;
                $costo = isset($_POST['costo']) ? $_POST['costo'] : NULL;

                $query = "UPDATE edizionecorso SET
                    numerominpartecipanti = '$minpar',
                    numeromaxpartecipanti = '$maxpar',
                    numerolezioni = '$nlezioni',
                    costo = '$costo'
                    WHERE codicecorso = '$codicecorso';";
                //var_dump($query);
                //update($query, $conn);
                 break;
            case 'persona' :
                $nome = isset($_POST['nomeis']) ? $_POST['nomeis'] : NULL;
                $cognome = isset($_POST['cognomeis']) ? $_POST['cognomeis'] : NULL;
                $numerotessera = isset($_POST['ntessera']) ? $_POST['ntessera'] : NULL;
                $datadinascita = isset($_POST['datais']) ? $_POST['datais'] : NULL;
                $genitore = isset($_POST['genitoreis']) ? $_POST['genitoreis'] : NULL;
                $nomepiscina = isset($_POST['nomepiscina']) ? $_POST['nomepiscina'] : NULL;
                $inizio= isset($_POST['inizio']) ? $_POST['inizio'] : NULL;
                $fine = isset($_POST['fine']) ? $_POST['fine'] : NULL;

                $query = "UPDATE iscritto SET
                    nome = '$nome',
                    cognome = '$cognome',
                    datadinascita = '$datadinascita',
                    genitore = '$genitore'
                    WHERE numerotessera = '$numerotessera';";

                $query .= "UPDATE tesseramento SET 
                    fine = '$fine'
                    WHERE inizio = '$inizio' AND nomepiscina = '$nomepiscina';";

                //var_dump($query);
                //update($query, $conn);
                break;
        }
        update($query, $conn);
    }
}



?>