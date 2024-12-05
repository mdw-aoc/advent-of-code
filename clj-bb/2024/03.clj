(def sample-input (slurp "03-sample.txt"))
(def actual-input (slurp "03-actual.txt"))

(defn parse-multiplications [input]
  (->> input
       (re-seq #"mul\((\d{1,3}),(\d{1,3})\)")
       (map rest) ; groups contain the entire match first, then the capture groups
       (flatten)
       (map #(Integer/parseInt %))
       (partition 2)
       (map #(apply * %))
       (reduce +)))

(println 161       (parse-multiplications sample-input))
(println 170068701 (parse-multiplications actual-input))

(def sample-input (slurp "03-sample2.txt")) ; redefine

(defn remove-donts [input]
  (let [dosplit   (str/split input #"do\(\)")
        dont      #"don't\(\)"
        dontsplit (map #(first (str/split % dont)) dosplit)]
    (apply str dontsplit)))

(println 48       (parse-multiplications (remove-donts sample-input)))
(println 78683433 (parse-multiplications (remove-donts actual-input)))
