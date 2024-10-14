<html>
    <head>
        <title>Visualizzazione</title>
        <style>
            table, th, td {
            text-align: left;
            border: 1px solid;
        }
        </style>
    </head>
    <body>
        <h1>Visualizzazione</h1>
        <div>
            <?php
            session_start();
            function check($query, $conn){
                $result = pg_query($conn,$query);
                if (!$result) {
                    echo "Si è verificato un errore.<br/>";
                    echo pg_last_error($conn);
                    exit();
                }
                return $result;
            }


            $conn = pg_connect("host=localhost port=5432 dbname=Piscine user=postgres password=unimi");
            if (!$conn) {//caso connessione fallita
              echo 'Connessione al database fallita.';
              exit();
            }

            $query = "SELECT nome FROM piscina";
            $piscine  = check($query, $conn);

            $query = "SELECT MIN(DATE_PART('year', inizio)) AS primocontratto, MAX(DATE_PART('year', inizio)) AS ultimocontratto
                        FROM contratto";
            $contratti = check($query, $conn);


            //form per la prima richiesta
            print("<form name =\"visualizza\" action =\"visualizza2.php\" method =\"POST\">
                <h2>Selezionare la struttura e l'anno di cui visualizzare il personale:</h2>");

            //select per la struttura
            print("<table><tr><th>Struttura:</th>");
            print("<td><select name = \"struttura\">");
            while ($row = pg_fetch_array($piscine)){
                print("<option>".$row['nome'] ."</option>");
            }
            print("</td></tr>");
           

            //select per l'anno
            print("<tr><th>Anno: </th>");
            print("<td><select name = \"anno\">");
            while ($row = pg_fetch_array($contratti)){
                $pc =  $row['primocontratto'];
                $uc =  $row['ultimocontratto']; 

            }
            for($i = $pc; $i <= $uc; $i++){
                print("<option>". $i ."</option>");
            }
            print("</select></td></tr>");
            print("</table>");
            print ("<br><input type=\"submit\" name=\"personale\" value=\"Visualizza\">");

            print("<br><hr>");


            //form per la seconda richiesta

            $query = "SELECT nome FROM piscina";
            $piscine  = check($query, $conn);

            print("<h2>Selezionare la struttura di cui visualizzare gli iscritti:</h2>");

            //select per la struttura
            print("<table><tr><th>Struttura:</th>");
            print("<td><select name = \"piscina\">");
            while ($row = pg_fetch_array($piscine)){
                print("<option>".$row['nome'] ."</option>");
            }
            print("</select></td></tr>");
            print("</table>");
            print ("<br><input type=\"submit\" name=\"iscritti\" value=\"Visualizza\">");
            print("</form>");
            ?>
        </div>
    </body>


</html>