import java.time.LocalDate;
import java.util.ArrayList;
import java.util.Scanner;

import javax.lang.model.util.ElementScanner14;

public class Main {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        ArrayList<Film> films = new ArrayList<Film>();

        Film f = null;

        System.out.println("Inserisci i film noleggiati nel formato `nome tipo data` (CTRL+D per terminare)");
        while(s.hasNext()){
            String titolo = s.next();
            String tipo = s.next();
            String[] dataTmp = s.next().split("-");
            int giorni = Integer.parseInt(dataTmp[0]);
            int mese = Integer.parseInt(dataTmp[1]);
            int anno = Integer.parseInt(dataTmp[2]);
            LocalDate l = LocalDate.of(anno,mese,giorni);

            if(tipo.equals("Azione"))
                f = new Azione(titolo, tipo, l);
            else if(tipo.equals("Commedia"))
                f = new Commedia(titolo, tipo, l);
            else
                f = new Dramma(titolo, tipo, l);

            films.add(f);
        }

        //System.out.println(films.size());

        double penaleTotale = 0.0;
        for (Film film : films) {
            penaleTotale += film.calcolaPenaleRitardo();
        }

        System.out.println("Penale totale: " + penaleTotale + "$");
    }
    
}
