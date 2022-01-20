create table personality(
  id serial primary key,
  name varchar,
  about varchar
);
INSERT INTO personality(name, about)
VALUES (
    'Eduardo Saverin',
    'Knowing how to take advantage of good opportunities is one of the main qualities of entrepreneur Eduardo Saverin, Brazilian by birth and currently a citizen of Singapore. 
    In 2004, while a student at Harvard, he helped fellow student Mark Zuckerberg create what would become Facebook.'
  ),
  (
    'Jorge Paulo Lemann',
    'Through the hands of Jorge Paulo Lemann and his inseparable partners Carlos Alberto Sicupira and Marcel Telles, Brazilian companies have become global. 
    Together with his partners, Lemann created Ambev, the result of the merger of rivals Brahma and Antárctica, and ended up forming the largest brewery in the world: ABInbev, which emerged from the merger of Ambev with the Belgian Interbrew and the acquisition of the American Anheuser-Busch. , owner of brands such as Budweiser and Stella Artois.'
  );