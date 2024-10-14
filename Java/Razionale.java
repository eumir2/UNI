public class Razionale{
//OVERVIEW: classe che rappresenta un numero razionale con il suo numeratore (intero) ed il suo denominatore (intero, diverso da 0).

    private int num;
    private int den;
    public static void main(String[] args) {
        int numeratore = Integer.parseInt(args[0]);
        int denominatore = Integer.parseInt(args[1]);

        try {
            Razionale r = new Razionale(numeratore, denominatore);
            System.out.println(r.toString());
        } catch (ArithmeticException e) {
            System.out.println(e.getMessage());
            System.exit(0);
        }

    }   

//CONSTRUCTORS
    public Razionale(int num, int den)throws ArithmeticException {
    //MODIFIES: this
    //EFFECTS: inizializza un'istanza del numero razionale. Se il denominatore  è zero
    //         lancia ArithmeticException  
        if(den != 0){
            this.num = num;
            this.den = den;
        }else 
            throw new ArithmeticException("Denominatore del razionale è zero");
    }
//METHODS
    public int getNum() {
        return num;
    }
    public void setNum(int num) {
        this.num = num;
    }
    public int getDen() {
        return den;
    }
    public void setDen(int den) {
        this.den = den;
    }

    public double valore(){
        return num/den;
    }

    @Override
    public String toString() {
        return "Il razionale è " + num + "/" + den + " o " + valore();
    }

    private boolean repOk(){
        if((Integer)num== null) return false;
        if((Integer)den == null) return false;

        return true;


    }




    




}