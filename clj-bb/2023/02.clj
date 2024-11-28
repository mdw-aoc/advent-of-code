#!/usr/bin/env bb

(defn parse-round [raw]
  (let [words  (-> raw str/trim (str/replace "," "") (str/split #"\s"))
        pairs  (->> words (partition 2) (map reverse) (map vec) vec)
        parsed (for [[key val] pairs] [(keyword key) (Integer/parseInt val)])]
   (into {} parsed)))

(defn parse-game [raw]
  (let [words  (str/split raw #"\s")
        id     (->> words second butlast (apply str))
        rounds (->> words (drop 2) (str/join " "))
        rounds (str/split rounds #";")
        rounds (map parse-round rounds)]
   {:id (Integer/parseInt id) :rounds rounds}))

(def maximums
  {:red   12
   :green 13
   :blue  14})

(defn is-valid-color? [[color-key color-count]]
  (or (nil? color-count) (<= color-count (color-key maximums))))

(defn round-is-valid? [round]
  (every? is-valid-color? round))

(defn game-is-valid? [game]
  (every? round-is-valid? (:rounds game)))

(defn part-1 [filename]
  (let [raw-games (-> filename slurp (str/split #"\n"))]
    (->> raw-games
         (map parse-game)
         (filter game-is-valid?)
         (map :id)
         (reduce +))))

(println "part 1 sample answer 8?   " (part-1 "02-sample.txt"))
(println "part 1 actual answer 2449?" (part-1 "02-actual.txt"))
