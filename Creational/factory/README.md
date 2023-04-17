<h1>工厂方法模式</h1>

<h2>工厂方法模式中的角色和职责</h2>

    简单工厂模式 + "开闭原则" = 工厂方法模式

&emsp;&emsp;抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。

&emsp;&emsp;工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。

&emsp;&emsp;抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。

&emsp;&emsp;具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象

<h2>工厂模式的优缺点</h2>

## 优点

&emsp;&emsp;不需要记住具体类名，甚至连具体参数都不用记忆。

&emsp;&emsp;实现了对象创建和使用的分离。

&emsp;&emsp;系统的可扩展性也就变得非常好，无需修改接口和原类。

&emsp;&emsp;对于新产品的创建，符合开闭原则。

## 缺点

&emsp;&emsp;增加系统中类的个数，复杂度和理解度增加。

&emsp;&emsp;增加了系统的抽象性和理解难度。

## 适用场景

&emsp;&emsp;客户端不知道它所需要的对象的类。

&emsp;&emsp;抽象工厂类通过其子类来指定创建哪个对象。