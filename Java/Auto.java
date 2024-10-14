import java.util.Scanner;

public class Auto {
//OVERVIEW: classe che implementa un'automobile con una certa capacità di serbatoio, velocità massima, e consumo medio (km per litro)  
    private static double capacitàSerbatoio;
    private static int velocitàMax;
    private static int consumoMedio;



    public static void main(String[] args) {
        double cS = Double.parseDouble(args[0]);
        int vM = Integer.parseInt(args[1]);
        int cM = Integer.parseInt(args[2]);

        int kmTot = 0;

        Scanner s = new Scanner(System.in);

        System.out.println("Inserisci una tratta (<L. riforniti> <km da fare> <velocità>)");
        while(s.hasNext()){
            System.out.println("Inserisci una tratta (<L. riforniti> <km da fare> <velocità>)");
            double lF = s.nextDouble();
            int km = s.nextInt();
            int v = s.nextInt();

            setCapacità(lF);
            kmTot += viaggia(km, v);
        }

        System.out.println("\nHai percorso " + kmTot + "km totali");
    }
    //CONSTRUCTORS
    public Auto (double cS, int vM, int cM){
    //REQUIRES: i parametri devono essere valori interi e positivi
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto macchina con la sua capacità, velocità e consumo
        capacitàSerbatoio = cS;
        velocitàMax = vM;
        consumoMedio = cM;
    }

    public double getCapacità(){
        return capacitàSerbatoio;
    }

    public int getVelocità(){
        return velocitàMax;
    }

    public int getConsumo(){
        return consumoMedio;
    }

    public void setVelocità(int v){
    //REQUIRES: il valore in input deve essere un intero positivo
    //MODIFIES: this
    //EFFECTS: modifica la velocità della macchina
        velocitàMax = v;
    }

    public static void setCapacità(double c){
    //REQUIRES: il valore in input deve essere intero positivo
    //MODIFIES: capacitàSerbatoio
    //EFFECTS: metodo che riempie il serbatoio con una certa quantità di carburante
     capacitàSerbatoio = c;
    }

    public static int viaggia(int d, int v){
    //REQUIRES: i valori in input devono essere interi e positivi
    //MODIFIES: capacitàSerbatoio
    //EFFECTS: pianificare un viaggio con in input: una distanza e una velocità di percorrenza. Il metodo deve calcolare il tempo di percorrenza ed aggiornare il carburante sottraendo quello consumato.
    //velocità = spazio/tempo => tempo = spazio/velocità
      double ore = d/v; //in ore
      int minuti = ((d/v) % 10 ) * 60;
      
      int consumoTmp = consumoMedio * d;
      if (consumoTmp > consumoMedio){
        System.out.println("Non hai carburante sufficiente");
        return 0;
      }else{
        capacitàSerbatoio -= consumoTmp;
        System.out.println("Tempo necessario: "+ ore + " ore " + minuti + " minuti");
        return d;
      }

    } 




}
