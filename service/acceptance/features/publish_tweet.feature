Feature: Publicar un tweet
  Como usuario de Tweeter
  Quiero poder publicar lo que pienso
  Para que el mundo sepa lo genial que soy

  Scenario: Publicación exitosa de un tweet
    Given que existe el usuario @womenwhogoba
    And que el usuario @womenwhogoba no ha twiteado nunca
    When el usuario @womenwhogoba envía el tweet con texto hola mundo
    Then en el timeline de @womenwhogoba aparece el tweet con texto hola mundo


