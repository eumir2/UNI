public interface Figura extends Comparable<Figura>{
//OVERVEW: interfaccia che modella una figura immutabile

    @Override
    default int compareTo(Figura o){
        if(this.perimetro() > o.perimetro())
            return -1;
        else if(this.perimetro() < o.perimetro())
            return 1;
        else
            return 0;
    }

    public double perimetro();

    public double area();

    
    
}
