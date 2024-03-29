## 解释器模式

### 定义

解释器模式(interpreter):给定一种语言，定义它的文法的一种表示，并定一个解释器，这个解释器使用该表示来解释语言中的句子。

解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。

对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以。

### 优点

1、可扩展性比较好，灵活;

2、增加了新的解释表达式的方式;

3、易于实现简单文法。

### 缺点

1、可利用场景比较少;

2、对于复杂的文法比较难维护;

3、解释器模式会引起类膨胀。

适用范围

解释器模式的代码实现比较灵活，没有固定的模板。

应用设计模式主要是应对代码的复杂性，实际上，解释器模式也不例外。

它的代码实现的核心思想，就是将语法解析的工作拆分到各个小类中，以此来避免大而全的解析类。

一般的做法是，将语法规则拆分成一些小的独立的单元，然后对每个单元进行解析，最终合并为对整个语法规则的解析。