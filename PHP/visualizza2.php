<?php
session_start();
    $htmlint = <<<NOW
    <HTML>
    <HEAD>
        <style>
        table, th, td {
            text-align: left;
            border: 1px solid;
            padding: 5px;
        }
        </style>
    </HEAD>
    <BODY> 
    NOW;
    function check($query, $conn){
        //var_dump($query);

        $result = pg_query($conn,$query);
        if (!$result) {
            echo "Si è verificato un errore.<br/>";
            echo pg_last_error($conn);
            exit();
        }
        return $result;
    };

    print($htmlint);
    $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
    if (!$conn) {//caso connessione fallita
    echo 'Connessione al database fallita.';
    exit();
    }


    //controllo quale operazione si deve svolgere
    if(isset($_POST['personale'])){

        $anno = isset($_POST['anno']) ? $_POST['anno'] : NULL;
        $nomepiscina = isset($_POST['struttura']) ? $_POST['struttura'] : NULL;
        $query = "SELECT i.nome, i.cognome, c.inizio, c.fine, c.tipo
                    from contratto as c JOIN istruttore as i ON c.badgeistruttore = i.numerobadge
                JOIN piscina as p ON p.nome = c.nomepiscina
                WHERE (DATE_PART('year', c.inizio) <= '$anno' AND c.tipo = 'Indeterminato')
                OR (DATE_PART('year', c.inizio) <= '$anno' AND DATE_PART('year', c.fine) >= '$anno' AND c.tipo = 'Determinato')
                AND c.nomepiscina = '$nomepiscina'";
        $contratti = check($query, $conn);

        print("<h1>Personale per la struttura '". $nomepiscina . "'</h2>");
        print("<table><tr>
                <th>Nome</td>
                <th>Cognome</th>
                <th>Inizio</th>
                <th>Fine</th>
                <th>Tipo</th></tr>");
        while ($row = pg_fetch_array($contratti)) {
            echo "<tr>";
            print("<td>".$row['nome']."</td>");
            print("<td>".$row['cognome']."</td>");
            print("<td>".$row['inizio']."</td>");
            print("<td>".$row['fine']."</td>");
            print("<td>".$row['tipo']."</td>");
        };       
        print("</tr></table>"); 
        echo ("<br><br>Operazione avvenuta con successo, <a href='visualizzazione.php'>RITORNA</a>");
    }else if(isset($_POST['iscritti'])){
        $piscina = isset($_POST['piscina']) ? $_POST['piscina'] : NULL;
        //var_dump($piscina);

        $query = "  SELECT DISTINCT(i.nome, i.cognome) 
                    FROM tesseramento as t JOIN iscritto as i ON t.numerotessera = i.numerotessera 
                    WHERE t.nomepiscina = '$piscina';";
        $iscritti  = check($query, $conn);

        

        if(pg_num_rows($iscritti) == 0){
            print("Non ci sono iscritti in questa piscina");
            echo ("<br><br>Se vuoi puoi <a href='visualizzazione.php'>riprovare</a>");
            exit();
        }

        print("<form name =\"iscritto\" action =\"iscritto.php\" method =\"POST\">
        <h2>Selezionare l'iscritto di cui si desidera visualizzare lo storico iscrizioni:</h2>");

        //select per la struttura
        print("<table><tr><th>Iscritto:</th>");
        print("<td><select name = \"iscritto\">");
        while ($row = pg_fetch_array($iscritti)){
            $datiIscritto = explode(",", $row['row']);
            print("<option>". $datiIscritto[0] . " ". $datiIscritto[1] . "</option>");
            //print("<input name=\"numerotessera\" type=\"hidden\" value=\"" . htmlspecialchars($row['numerotessera']) . "\">");
        }
        print("</select></td></tr>");
        print("<input name=\"piscina\" type=\"hidden\" value=\"" . htmlspecialchars($_POST['piscina']) . "\">");
        print("</table>");

        print ("<br><input type=\"submit\" name=\"visualize\" value=\"Visualizza\">");
        print("</form>");

    }






?>