(def sample-input (slurp "03-sample.txt"))
(def actual-input (slurp "03-actual.txt"))

(def expression-pattern #"mul\((\d{1,3}),(\d{1,3})\)")
(defn parse-multiplications [input]
  (->> input
       (re-seq expression-pattern)
       (map rest) ; groups contain the entire match first, then the capture groups
       (flatten)
       (map #(Integer/parseInt %))
       (partition 2)
       (map #(apply * %))
       (reduce +)))

(println 161       (parse-multiplications sample-input))
(println 170068701 (parse-multiplications actual-input))

(def sample-input (slurp "03-sample2.txt")) ; redefine

(def do-pattern   #"do\(\)")
(def dont-pattern #"don't\(\)")
(defn remove-donts [input]
  (let [do-split   (str/split input do-pattern)
        dont-split (map #(first (str/split % dont-pattern)) do-split)]
    (apply str dont-split)))

(println 48       (parse-multiplications (remove-donts sample-input)))
(println 78683433 (parse-multiplications (remove-donts actual-input)))
