import java.util.ArrayList;
import java.util.InputMismatchException;
import java.util.Scanner;

public class Punto {
//OVERVIEW: classe che rappresenta il punto nel piano cartesiano

    private double x;
    private double y;

public static void main(String[] args) {
    Scanner s = new Scanner(System.in);
    ArrayList<Punto>  punti = new ArrayList<Punto>();

    System.out.println("Inserire i punti (<x> <y>) \nTerminare la lettura con CTRL+D");

    while(s.hasNext()){
        try {
            String tmpX = s.next();
            String tmpY = s.next();
            double x = 0.0;
            double y = 0.0;
            if(tmpX.equals("0"))
                x = 0.0;
            else
                x = Double.parseDouble(tmpX);
            if(tmpY.equals("0"))
                y = 0.0;    
            else
                y = Double.parseDouble(tmpY);
            
            Punto tmp = new Punto(x,y);
            //System.out.println(tmp.toString());
            punti.add(tmp);

        } catch (InputMismatchException e) {
            System.out.println("Input inserito non valido");
            break;
        }
        
    }

    //System.out.println(punti);

    double maxDist = 0.0;
    Punto p1 = new Punto();
    Punto p2 = new Punto();
    for (Punto punto1 : punti) {
        for (Punto punto2 : punti) {
            double distanzaTmp = punto1.distanza(punto2);
            if (distanzaTmp > maxDist){
                maxDist = distanzaTmp;
                p1.setX(punto1.x);
                p1.setY(punto1.y);

                p2.setX(punto2.x);
                p2.setY(punto2.y);

                //System.out.println(punto1.toString() + " " + punto2.toString() + "\t" + distanzaTmp );
            }
        }
    }

    System.out.println("I punti più distanti sono:");
    System.out.println(p1.toString());
    System.out.println(p2.toString());
}

//CONSTRUCTORS
    public Punto(double x, double y){
    //MODIFIES: this
    //EFFECTS: istanzia un nuovo oggetto Punto
        this.x = x;
        this.y = y;
    }

    public Punto(){
    //MODIFIES: this
    //EFFECTS: istanzia un nuovo oggetto Punto con coordinate nell'origine  
        this.x = 0;
        this.y = 0;
    }


//METHODS

    public double getX(){
        return this.x;
    }

    public double getY(){
        return this.y;
    }


    public void setX(double x){
        this.x = x;
    }

    public void setY(double y){
        this.y = y;
    }

    public void setXY(double x, double y){
        this.x = x;
        this.y = y;
    }

    public String toString(){
        return "{Punto: " + this.x + ", " + this.y + "}";
    }

    public double distanza(Punto p){
    //REQUIRES: p deve essere un punto valido
    //EFFECTS: restituisce la distanza euclidea fra il punto p e il punto su cui il metodo viene invocato.
       double xTmp = Math.pow((p.x - this.x), 2);
       double yTmp = Math.pow((p.y - this.y), 2);
       return Math.sqrt(xTmp + yTmp);  
    }

    @Override
    public boolean equals(Object obj) {

        if(obj == null) 
            return false;
        if (!(obj instanceof Punto)) 
            return false;
        
        Punto tmp = (Punto)obj;
        if((tmp.x != this.x) || (tmp.y != this.y))
             return false;
        
        return true;
    }
}
