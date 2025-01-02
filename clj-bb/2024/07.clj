(defn parse-line [line]
  (as-> line $
        (str/replace $ #":" "")
        (str/split $ #"\s")
        (map clojure.edn/read-string $)))

(defn parse-input [input]
  (as-> input $
        (str/split $ #"\n")
        (map parse-line $)))

(defn operators [options n]
  (lazy-seq
    (if (zero? n)
      [[]]
      (for [option options
            combos (operators options (dec n))]
        (cons option combos)))))

(defn calculate [operands operators]
  (loop [total     (first operands)
         operands  (rest operands)
         operators operators]
    (if (empty? operators)
      total
      (recur ((first operators) total (first operands))
             (rest operands)
             (rest operators)))))

(defn possible? [operator-combos [target & operands]]
  (let [combos (nth operator-combos (dec (count operands)))]
    (->> combos
         (map (partial calculate operands))
         (some #(= target %)))))

(defn calibration-result [filename operator-combos]
  (->> filename
       slurp
       parse-input
       (filter (partial possible? operator-combos))
       (map first)
       (reduce +)))
 
(def part1-operator-combos
  (->> (range) (take 12) (map (partial operators [+ *])) vec))

(println 3749          (calibration-result "07-sample.txt" part1-operator-combos))
(println 5702958180383 (calibration-result "07-actual.txt" part1-operator-combos))

(defn || [a b]
  (clojure.edn/read-string (str a b)))

(def part2-operator-combos
  (->> (range) (take 12) (map (partial operators [+ * ||])) vec))

(println 11387          (calibration-result "07-sample.txt" part2-operator-combos))
(println "Please stand by, the answer to part 2 with the actual input takes about 30 seconds... (yuck)")
(println 92612386119138 (calibration-result "07-actual.txt" part2-operator-combos))

