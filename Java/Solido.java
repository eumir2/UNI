abstract class Solido implements Comparable<Solido>{
    private String tipo;
    private double volume;

//CONSTRUCTOR
    public Solido(String tipo){
    //MODIFIES: this
    //EFFECTS: istanzia un Solido con il proprio tipo
        this.tipo = tipo;
        this.volume = 0.0;
    }

    abstract double volume();

    public String tipo(){
        return this.tipo;
    }

    @Override
    public abstract int compareTo(Solido o);


}
