# Linked List Go

Односвязный список на Go с ограничением максимальной длины до константного количества узлов (10). 

Этот проект представляет собой реализацию односвязного списка на Go. Односвязный список — это структура данных, состоящая из узлов, где каждый узел содержит данные и ссылку на следующий узел в последовательности. Это позволяет эффективно добавлять и удалять элементы в списке.

Произведена декомпозиция проекта с учетом инкапусляции. Проект содержит следующие каталоги:

1. companyOps
- производит преобразование строки с названиями компаний в мапу, содержащую ID компании и её название;
- присваивает узлу значение рандомного элемента из мапы, а затем удаляет этот элемент для исключения повторного использования.

2. socketMaker
- создает сокет при помощи функций формирования рандомного IP и порта.

3. nodeAndListOps
- предлагает ввести количество элементов для создания односвязного списка или прибавления к уже созданному;
- создает узлы, содержащие инфу о названии компании, ID и сокете;
- создает связаный список путем добавления элемента в начало списка;
- добавляет узлы, не превышая константной длины списка (10) и удаляя самые ранние узлы по принципу стека в случае превышения общего количество узлов в списке;
- в случае превышения константной длины списка количеством добавляемых узлов возвразщает новый список константной длины;
- выводит количество узлов списка;
- выводит информацию о всех узлах списка.

4. errorOps
- в случае превышения вводимым количеством узлов длины мапы логирует, выдает ошибку и прекращает работу программы.
