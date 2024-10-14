import java.util.Scanner;

public class Tris {
    // OVERVIEW: classe che modella il gioco del tris con due giocatori

    private char[][] tabellone;
    private boolean giocatore; // O = true , X = false
    private int mosse;

    // CONSTRUCTORS
    public Tris() {
        // MODIFIES: this
        // EFFECTS: crea un nuovo gioco col tabellone vuoto e assegna il turno al
        // giocatore `O`
        this.tabellone = new char[3][3];
        this.giocatore = true;
        this.mosse = 0;
    }

    // METHODS
    public char turno() {
        if (this.giocatore)
            return 'O';
        else
            return 'X';
    }

    public void mossa(int x, int y) throws IllegalArgumentException {
        // MODIFIES: this.tabellone
        // EFFECTS: esegue la mossa del giocatore, se la cella è piena, lancia
        // IllegalArgumentException
        if (this.tabellone[x - 1][y - 1] == ' ')
            throw new IllegalArgumentException("Cella già occupata");
        else {
            if (!(terminato()))
                this.mosse++;

            if (this.giocatore)
                this.tabellone[x - 1][y - 1] = 'O';
            else
                this.tabellone[x - 1][y - 1] = 'X';
        }

        // cambio turno
        this.giocatore = !this.giocatore;
    }

    public boolean vittoria() {
        // EFFECTS: restituisce true se il giocatore di turno ha vinto (tre dei suoi
        // simboli in fila su riga, colonna o diagonale)
        boolean flag = false;

        // controllo le righe
        for (int i = 0; i < 3; i++) {
            if (this.tabellone[i][0] == this.tabellone[i][1] && this.tabellone[i][0] == this.tabellone[i][2])
                flag = true;
        }
        // controllo le colonne
        for (int i = 0; i < 3; i++) {
            if (this.tabellone[0][i] == this.tabellone[1][i] && this.tabellone[0][i] == this.tabellone[2][i])
                flag = true;
        }

        // controllo la diagonale principale
        if (this.tabellone[0][0] == this.tabellone[1][1] && this.tabellone[0][0] == this.tabellone[2][2])
            flag = true;

        // controllo la diagonale secondaria
        if (this.tabellone[0][2] == this.tabellone[1][1] && this.tabellone[0][2] == this.tabellone[2][0])
            flag = true;

        return flag;
    }

    public boolean terminato() {
        if (this.mosse == 9)
            return true;
        else
            return false;
    }

    public String toString() {
        String tmp = "";
        tmp += "----------------- \n |R/C| 1 | 2 | 3 | \n -----------------\n";
        for (int i = 0; i <= 2; i++) {
            tmp += " | " + (i + 1) + " | ";
            for (int j = 0; j < 3; j++) {
                if(tabellone[i][j] == ' ')
                    tmp += String.format("|%3s|", "  |");
                else    
                    tmp += tabellone[i][j] + " | ";
            }
            tmp += "\n-----------------\n";
        }
        return tmp;
    }

    // MAIN
    public static void main(String[] args) {
        Tris t = new Tris();

        Scanner s = new Scanner(System.in);
        System.out.println("Mossa di " + t.turno());
        System.out.println("Inserisci: x y");
        while (s.hasNext()) {
            int x = s.nextInt();
            int y = s.nextInt();

            try {
                t.mossa(x, y);
            } catch (IllegalArgumentException e) {
                System.out.println(e.getMessage());
            }

            System.out.println(t.toString());
            if(t.vittoria() && t.mosse >= 5){

                if(t.giocatore)
                    System.out.println("Ha vinto X");
                else    
                    System.out.println("Ha vinto O");
                
                System.exit(0);
            }

            System.out.println("Mossa di " + t.turno());
            System.out.println("Inserisci: x y");
           
        }
        s.close();

    }
}
