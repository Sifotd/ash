public class maopao {
    public static void main(String[] args) {
        int []a={78,85,32,64,91,46,77};
        int temp=0;
        for(int i=0;i<7;i++)
        {
            for(int j=0;j<6-i;j++)
            {
                if(a[j]>a[j+1]){
                 temp=a[j];
                a[j]=a[j+1];
                a[j+1]=temp;
                        }

            }
        }
        for(int i=0;i<7;i++)
        {
            System.out.println(a[i]);
        }
    }
}
