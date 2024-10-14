<!--basic_ins.php-->
<?php
session_start();
$htmlint = <<<NOW
<HTML>
  <HEAD>
    <style>
      table, th, td {
        text-align: left;
        border: 1px solid;
      }
    </style>
  </HEAD>
<BODY> 
NOW;
//print_R($_GET);
//print_R($_SESSION);
//print_R($_POST);


//funzione che tramite una transazione effettua le query
function commit($conn, $query){
    //var_dump($query);

   pg_query("BEGIN");
    $result = pg_query($conn, $query);
    if (!$result) {
        echo "Si è verificato un errore.<br/>";
        echo pg_last_error($conn);
        pg_query("ROLLBACK");
        echo ("<br><br>Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
        exit();
    }else{
        pg_query("COMMIT");
        echo("Inserimento avvenuto con successo");
        echo "Torna alla <a href='inserimento_e_modifica.php'>home</a>";
    }
}

if (isset($_POST['toinsert']) && isset($_SESSION['tbl'])) {//sono stati passati correttamente i dati
  $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
  if (!$conn) {
    echo 'Connessione al database fallita.';
    exit();
  }
  else {
        $query = "";
        $tbl = $_SESSION['tbl'];
        switch($tbl){
            case 'piscina' :
                //echo "Connessione riuscita."."<br/>";
                //print_r($_POST);
                //attributi piscina
                $tel = isset($_POST['telc']) ? $_POST['telc'] : NULL;
                //testo anche la composizione
                $pattern = '/^[0-9]+$/';
                $tel = preg_match($pattern, $tel) ? $tel : 'not valid';
                $nomepiscina = isset($_POST['nome']) ? $_POST['nome'] : NULL;
                $via = isset($_POST['via']) ? $_POST['via'] : NULL;
                $nciv = (int)(isset($_POST['nciv']) and is_numeric($_POST['nciv'])) ? $_POST['nciv'] : NULL;
                $nvasche = (int)(isset($_POST['nvasche']) and is_numeric($_POST['nvasche'])) ? $_POST['nvasche'] : NULL;
                $inizio = isset($_POST['inizio']) ? $_POST['inizio'] : NULL;
                $fine = isset($_POST['fine']) ? $_POST['fine'] : NULL;

                //attributi responsabile
                $cf = isset($_POST['CF']) ? $_POST['CF'] : NULL; 
                $nomer = isset($_POST['nomer']) ? $_POST['nomer'] : NULL;
                $cognomer = isset($_POST['cognomer']) ? $_POST['cognomer'] : NULL;
                $datar = isset($_POST['datar']) ? $_POST['datar'] : NULL;
                $telr = isset($_POST['telr']) ? $_POST['telr'] : NULL;

                $diasettimana = array("Lunedi", "Martedi", "Mercoledi", "Giovedi", "Venerdi", "Sabato", "Domenica");
                $reperibilita = array();
                foreach ($diasettimana as $giorno) {
                    if(isset($_POST[$giorno])){
                        array_push($reperibilita, $giorno);
                    }
                }
                

                //attributi vasche
                $ncors = isset($_POST['ncors']) ? $_POST['ncors'] : NULL;
                $tipologia = isset($_POST['tipologia']) ? $_POST['tipologia'] : NULL;
                $iniziof = isset($_POST['iniziof']) ? $_POST['iniziof'] : NULL;
                $finef = isset($_POST['finef']) ? $_POST['finef'] : NULL;

                //var_dump($ncors);


                $query .= "INSERT INTO responsabile (CF, nome, cognome, datadinascita)
                VALUES ('$cf', '$nomer', '$cognomer', '$datar'); ";
                //echo $query;

                $tmp = "INSERT INTO " . $tbl . " (nome, cfresponsabile, via, n_civ , apertura, chiusura, nvasche) 
                            VALUES ('$nomepiscina','$cf','$via','$nciv','$inizio','$fine','$nvasche'); ";

                //var_dump($tel);
                if($tel!='not valid'){
                    $query .= $tmp;
                }

                $query .= "INSERT INTO Telefono (Numero, TesseraIscritto, BadgeIstruttore, NomePiscina, CFResponsabile, CodiceMedico)
                            VALUES ('$tel', NULL, NULL, '$nomepiscina', NULL, NULL); ";

                $query .= "INSERT INTO Telefono (Numero, TesseraIscritto, BadgeIstruttore, NomePiscina, CFResponsabile, CodiceMedico)
                            VALUES ('$telr', NULL, NULL, NULL, '$cf', NULL); ";

                $diasettimana = array("Lunedi", "Martedi", "Mercoledi", "Giovedi", "Venerdi", "Sabato", "Domenica");
                foreach ($diasettimana as $giorno) {
                    if(isset($_POST[$giorno])){
                        $query .= "INSERT INTO gestione (cfresponsabile, nomepiscina, giorno)
                                    VALUES ('$cf', '$nomepiscina', '$giorno'); " ;
                    }
                }
                
                    /*else {
                        //$cmdtuples = pg_affected_rows($result);
                        echo "Inserimento avvenuto con successo<br><a href='select_basic.php'>ritorna</a>";
                    }*/

                for($j = 0; $j < (int)$_POST['nvasche']; $j++){
                    $query .= "INSERT INTO Vasca (id, tipo, NomePiscina, NumeroCorsie, InizioFruizione, FineFruizione)
                    VALUES ('".($j + 1)."', '$tipologia[$j]', '$nomepiscina', '$ncors[$j]', '$iniziof[$j]', '$finef[$j]'); ";
                }

            break;
            case 'istruttore' :
                $tel = isset($_POST['teli']) ? $_POST['teli'] : NULL;
                //testo anche la composizione
                $pattern = '/^[0-9]+$/';
                $tel = preg_match($pattern, $tel) ? $tel : 'not valid';
                $nome = isset($_POST['nomei']) ? $_POST['nomei'] : NULL;
                $cognome = isset($_POST['cognomei']) ? $_POST['cognomei'] : NULL;
                $nascita = (isset($_POST['datai'])) ? $_POST['datai'] : NULL;
                /*$nuoto = isset($_POST['nuoto']) ? $_POST['nuoto'] : NULL;
                $fitness = isset($_POST['fitness']) ? $_POST['fitness'] : NULL;
                $pilates = isset($_POST['pilates']) ? $_POST['pilates'] : NULL;*/
                
                //ottengo le qualifiche
                $qual = "SELECT * FROM qualifica"; 
                $qualifiche = pg_query($conn, $qual);

                $tmp = array();
                while ($row = pg_fetch_array($qualifiche)){
                    if(isset($_POST[$row['nome']])){
                        array_push($tmp, $row['nome']);
                    }
                }
                

                $contratto = isset($_POST['contratto']) ? $_POST['contratto'] : NULL;
                $inizio = isset($_POST['inizioi']) ? $_POST['inizioi'] : NULL;
                $ferie= isset($_POST['ferie']) ? $_POST['ferie'] : NULL;
                $nomepiscina = isset($_POST['nomep']) ? $_POST['nomep'] : NULL;


                $badgetmp =  pg_query($conn, "SELECT * FROM istruttore");
                $result = pg_num_rows($badgetmp);
                if (!$result) {
                    echo "Si è verificato un errore.<br/>";
                    echo pg_last_error($conn);
                    exit();
                }
                else{
                    $numerobadge = $result+1;
                    $query .= "INSERT INTO istruttore (numerobadge, nome, cognome, datadinascita)
                                VALUES ('$numerobadge','$nome', '$cognome', '$nascita'); " ;
                } 
                

                //echo($numerobadge);
                //inserisco nella tabella possiede
                foreach($tmp as $q){
                    $query .= "INSERT INTO possiede (badgeistruttore, qualifica)
                    VALUES ('$numerobadge', '$q'); ";
                }
                    
                $query .= "INSERT INTO Telefono (Numero, TesseraIscritto, BadgeIstruttore, NomePiscina, CFResponsabile, CodiceMedico)
                            VALUES ('$tel', NULL, '$numerobadge', NULL, NULL, NULL); ";

                $tmpd = str_replace('/','-', $inizio);
                $time = strtotime($tmpd);
                switch($contratto){
                    case 'Indeterminato' :
                        $tipo = 'Indeterminato';
                        break;
                    case 'Mensile';
                        $durata = 1;
                        $tipo = 'Determinato';
                        $fine = date('Y-m-d', strtotime('+'. $durata .'months', $time));
                        break;
                    case 'Bi-mestrale';
                        $durata = 2;
                        $tipo = 'Determinato';
                        $fine = date('Y-m-d', strtotime('+'. $durata .'months', $time));
                        break;
                    case 'Tri-mestrale';
                        $durata = 3;
                        $tipo = 'Determinato';
                        $fine = date('Y-m-d', strtotime('+'. $durata .'months', $time));
                        break;
                    case 'Se-mestrale';
                        $durata = 6;
                        $tipo = 'Determinato';
                        $fine = date('Y-m-d', strtotime('+'. $durata .'months', $time));
                        break;
                };

                if($tipo == 'Indeterminato'){
                    $query .= "INSERT INTO contratto(badgeistruttore, nomepiscina, inizio, fine, durata, tipo, ferie)
                            VALUES('$numerobadge', '$nomepiscina', '$inizio', NULL , NULL, '$tipo', '$ferie'); ";
                }else{
                    $query .= "INSERT INTO contratto(badgeistruttore, nomepiscina, inizio, fine, durata, tipo, ferie)
                            VALUES('$numerobadge', '$nomepiscina', '$inizio', '$fine', '$durata', '$tipo', '$ferie'); ";
                }


            break;
            case 'corso' : 
                $nomecorso = isset($_POST['nomec']) ? $_POST['nomec'] : NULL;
                $nomepiscina = isset($_POST['nomep']) ? $_POST['nomep'] : NULL;
                $tipologia = isset($_POST['tipologiac']) ? $_POST['tipologiac'] : NULL;
                $livello = isset($_POST['livelloc']) ? $_POST['livelloc'] : NULL;

                //controllo che nella piscina ci siano delle vasche che possono ospitare il corso
                $check = "SELECT v.id
                FROM vasca AS v
                JOIN piscina AS p ON v.nomepiscina = p.nome
                WHERE p.nome = '$nomepiscina' AND v.tipo = '$tipologia'";

                $result = pg_query($conn, $check);
                if (!$result) {
                    echo "Si è verificato un errore.<br/>";
                    echo pg_last_error($conn);
                    exit();
                }else if (pg_num_rows($result) == 0){
                    echo("Non ci sono vasche idonee a questo corso ");
                    echo ("Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
                    //exit();
                }else{
                    //inserimento in corso
                    $query .= "INSERT INTO corso (nomecorso, tipologia, livello)
                    VALUES('$nomecorso', '$tipologia', '$livello');";

                    //inserimento in offre
                    $query .= "INSERT INTO offre(nomecorso, livellocorso, nomepiscina)
                                VALUES('$nomecorso', '$livello', '$nomepiscina');";
                } 
            break;
            case 'persona' :
                $nome = isset($_POST['nomeis']) ? $_POST['nomeis'] : NULL;
                $cognome = isset($_POST['cognomeis']) ? $_POST['cognomeis'] : NULL;
                $numerotessera = isset($_POST['ntessera']) ? $_POST['ntessera'] : NULL;
                $pattern = '/^[0-9]+$/';
                $numerotessera = preg_match($pattern, $numerotessera) ? $numerotessera : 'not valid';
                $datadinascita = isset($_POST['datais']) ? $_POST['datais'] : NULL;
                $genitore = isset($_POST['genitoreis']) ? $_POST['genitoreis'] : NULL;
                $nomepiscina = isset($_POST['piscina']) ? $_POST['piscina'] : NULL;
                $ii = isset($_POST['datainizio']) ? $_POST['datainizio'] : NULL;
                //var_dump($edizionecorso);

                //informazioni medico e certificato medico
                $generalita_medico = isset($_POST['medico']) ? $_POST['medico'] : NULL;
                $medico  = explode(" ",$generalita_medico); //posizione 0 = nome, posizione 1 = cognome

                $t = isset($_POST['datarilascio']) ? $_POST['datarilascio'] : NULL;

                //calcolo la data di fine iscrizione
                $tmpd = str_replace('/','-', $ii);
                $time = strtotime($tmpd);
                $inizioiscrizione = getdate($time);
                $fineiscrizione = date('Y-m-d', strtotime('+12 months', $time));
                //var_dump($inizioiscrizione);

                //calcolo la data di scadenza
                $tmpd = str_replace('/','-', $t);
                $time = strtotime($tmpd);
                $datarilascio = getdate($time);
                $datascadenza = date('Y-m-d', strtotime('+12 months', $time));
                ///var($inizioiscrizione);

                $timestampInizioIscrizione = strtotime($inizioiscrizione['year'] . '-' . $inizioiscrizione['mon'] . '-' . $inizioiscrizione['mday']);
                $timestampDataRilascio = strtotime($datarilascio['year'] . '-' . $datarilascio['mon'] . '-' . $datarilascio['mday']);

                if ($timestampDataRilascio > $timestampInizioIscrizione) {
                    echo "La data di rilascio è successiva alla data di inizio iscrizione.";
                    echo ("Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>");
                    exit();
                }


                //var_dump($_POST);
                $medico = "SELECT codiceregionaleasl FROM medico WHERE nome = '$medico[0]' AND cognome = '$medico[1]'";
                $codiceregionale = pg_query($conn, $medico);
                if (!$codiceregionale) {
                echo "Si è verificato un errore.<br/>";
                echo pg_last_error($conn);
                exit();
                } 

                while ($row = pg_fetch_array($codiceregionale)){
                    $codicer = $row['codiceregionaleasl'];
                }

                $query .= "INSERT INTO iscritto(numerotessera, nome, cognome, datadinascita, genitore)
                            VALUES('$numerotessera', '$nome', '$cognome', '$datadinascita', '$genitore'); ";
                
                $query .=  "INSERT INTO certificatomedico(numerotessera, datarilascio, codicemedico, datascadenza)
                            VALUES('$numerotessera', '$t', '$codicer', '$datascadenza');";

                $query .= "INSERT INTO tesseramento(numerotessera, inizio, nomepiscina, fine)
                            VALUES('$numerotessera', '$ii', '$nomepiscina', '$fineiscrizione'); ";
        } 

        //eseguo la query creata
        commit($conn, $query);
    }
}
else {
    echo "I dati passati non sono conformi alle richieste.<br>";
    echo "Se vuoi puoi <a href='inserimento_e_modifica.php'>riprovare</a>";
}
?>
</BODY>
</HTML>