(def sample-input (slurp "05-sample.txt"))
(def actual-input (slurp "05-actual.txt"))

(defn parse-input [content]
  (let [parts (str/split content #"\n\n")]
    {:rules   (set (map #(str/split % #"\|") (str/split (first parts) #"\n")))
     :updates (map #(str/split % #",") (str/split (second parts) #"\n"))}))

(defn is-valid? [rules update]
  (->> update
       (partition 2 1)
       (every? #(contains? rules %))))

(defn part1 [input]
  (let [parsed (parse-input input)]
    (->> parsed
         :updates
         (filter (partial is-valid? (parsed :rules)))
         (map #(nth %1 (/ (count %1) 2)))
         (map #(Integer/parseInt %))
         (reduce +))))

(println 143  (part1 sample-input))
(println 4814 (part1 actual-input))
