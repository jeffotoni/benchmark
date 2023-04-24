#rust hyper v4

Esta é a versão mais lenta, esta verssão removemos o Arc<ReqwestClient> e a criação de um novo ReqwestClient para cada chamada, isto resultará em um desempenho mais lento. Porque não estamos utilizando do pool de conexão e da reutilização de conexões oferecidos pelo ReqwestClient quando ele é compartilhado entre várias chamadas.

Ao usar Arc<ReqwestClient>, você permite que o cliente seja compartilhado entre várias chamadas sem precisar criar uma nova instância a cada vez. Isso permite que o ReqwestClient gerencie um pool de conexões e as reutilize, melhorando significativamente o desempenho.