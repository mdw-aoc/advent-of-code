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

(defn swap-items [lst idx1 idx2]
  (let [vector (vec lst)]
    (assoc vector
           idx1 (vector idx2)
           idx2 (vector idx1))))

(defn resort [rules update]
  (loop [i 2 update update]
    (if (is-valid? rules update)
      update
      (let [part (take i update)]
        (if (is-valid? rules part)
          (recur (inc i) update)
          (recur 2 (swap-items update (- i 2) (- i 1))))))))

(defn part2 [input]
  (let [parsed  (parse-input input)
        rules   (parsed :rules)
        updates (parsed :updates)]
    (->> updates
         (remove (partial is-valid? rules))
         (map (partial resort rules))
         (map #(nth %1 (/ (count %1) 2)))
         (map #(Integer/parseInt %))
         (reduce +))))

(println 123  (part2 sample-input))
(println 5448 (part2 actual-input))