(require '[clojure.xml :as xml])

; XML structure looks like this:
; <record>
;  <field name="Country or Area" key="ABW">Aruba</field>
;    <field name="Item" key="SP.POP.TOTL">Population, total</field>
;    <field name="Year">1960</field>
;    <field name="Value">54208</field>
; </record>
; And all we want is the population: 54208
; We can find the population by choosing nodes with name="Value"
; and discarding all other tags.

(defn populations []
  (remove nil?
    (for [node (xml-seq (xml/parse (clojure.java.io/file "data.xml")))
      :when (= (get-in node [:attrs :name]) "Value")]
     (first (:content node)))))

(let [countries (count (populations))]
  (print (map #(float (/ % countries)) (vals (frequencies (map #(first %) (populations)))))))
