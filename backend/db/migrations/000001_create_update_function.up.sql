CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   -- updated_atが明示的に変更されていない場合のみ、NOW()を設定
   IF NEW.updated_at = OLD.updated_at THEN
      NEW.updated_at = NOW();
   END IF;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';