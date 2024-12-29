(defn parse-line [line]
  (as-> line $
        (str/replace $ #":" "")
        (str/split $ #"\s")
        (map clojure.edn/read-string $)))

(defn parse-input [input]
  (as-> input $
        (str/split $ #"\n")
        (map parse-line $)))

(defn possible? [[target current & operands]]
  (and (not (> current target))
       (or (and (empty? operands) (= current target))
           (or (possible? (concat [target (+ current (or (first operands) 0))] (rest operands)))
               (possible? (concat [target (* current (or (first operands) 0))] (rest operands)))))))

(defn part1 [filename]
  (->> filename
       slurp
       parse-input
       (filter possible?)
       count))

(println 3 (part1 "07-sample.txt"))
(println 3 (part1 "07-actual.txt")) ; 318 is too low
