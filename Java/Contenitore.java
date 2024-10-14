package PROVAESAME;

public abstract class Contenitore implements Comparable<Contenitore>{
//OVERVIEW: classe che modella un contenitore generico, le istanze saranno mutabili perchè
//          il riempimento del contenitore può cambiare
    private double  volume;
    private double liquido;
    private String nome;

//CONSTRUCTORS
    public Contenitore(double liquido, String nome)throws ExceededCapacityException{
    //MODIFIES: this
    //EFFECTS: istanzia un contenitore che contiene un volume di liquido pari a "liquido" e il nome del liquido.
    //         se Se si cerca di creare un Contenitore con più liquido della sua capienza, viene lanciata una ExceededCapacityException
        if(this.volume < liquido)
            throw new ExceededCapacityException("ma il liquido ha un volume di: " + liquido);
        else{
            this.liquido = liquido;
            this.nome = nome;
        }
    }


//METHODS
    public double getVolume() {
        return volume;
    }


    public double getLiquido() {
        return liquido;
    }


    public String getNome() {
        return nome;
    }

    public void setLiquido(double liquido) {
        this.liquido = liquido;
    }


    public void setNome(String nome) {
        this.nome = nome;
    }

    protected void setVolume(double d){
        this.volume = this.calcolaVolume();
    }
    

    abstract double calcolaVolume();
    //MODFIES: this.volume
    //EFFECTS: calcola il volume del Contenitore

    @Override
    public int compareTo(Contenitore o) {
    //EFFECTS: comparatore inverso, restituisce 1 se this è minore di o
    //         restituisice -1 se this è maggiore di o
        if(this.getVolume() < o.getVolume())
            return 1;
        else if(this.getVolume() > o.getVolume())
            return -1;
        else
            return 0;
    }

    public void versaLiquido(Contenitore o) throws IncompatibleLiquidsException{
    //MODIFIES: this.liquido, o
    //EFFECTS: versa il liquido del contenitore o nel Contenitore this, se i liquidi non sono compatibili lancia IncompatibleLiquidsException
        if(this.getNome().equals(o.getNome())){
            if(this.getLiquido() + o.getLiquido() <= this.getVolume()){
                this.setLiquido(this.getLiquido() + o.getLiquido());
                o.setLiquido(0.0);
                o.setNome("");
            }else if(this.getLiquido() + o.getLiquido() > this.getVolume()){
                o.setLiquido(o.getLiquido() - (this.getVolume() - this.getLiquido()));
                this.setLiquido(this.getVolume()); 
            }   
        }else{
            throw new IncompatibleLiquidsException("Liquido " + this.getLiquido() + "non compatibile con " + o.getLiquido());

        }
    }

    @Override
    public String toString() {
        String tmp = "\t(capienza: " + this.getVolume() + ", contenuto: " + this.getLiquido() + ", liquido: " + this.liquido;
        return tmp += ")";
    }
   
}
